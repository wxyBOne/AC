<template>
    <div class="publish-container">
        <div id="publish-stars" ref="publishStarsContainer"></div>
        <img src="../img/logo.png" alt="Logo" class="publish-logo" />
        <div class="publish-room-out">
            <div class="publish-content">
                <div class="publish-tit">
                    <div class="publish-title-row">
                        Register and Publish Your Digital Artwork
                    </div>
                    <div class="publish-subtitle-row">
                        Share Your Creativity on ArtChain.
                    </div>
                </div>

                <div class="publish-description-row">
                    <div class="input-group-p">
                        <label for="title">艺术品标题*</label>
                        <input type="text" id="title" placeholder="输入标题" v-model="title">
                    </div>
                    <div class="input-group-p">
                        <label for="description">艺术品描述*</label>
                        <textarea id="description" placeholder="输入描述" v-model="description"></textarea>
                    </div>
                    <div class="input-group-p">
                        <label for="title">艺术品价格*</label>
                        <input type="text" id="title" placeholder="输入价格" v-model="price">
                    </div>

                    <div class="Art-upload-container">
                        <div class="Art-image-tit">艺术品图片*</div>
                        <p class="Art-image-tip">限制上传一张图片</p>
                        <div class="Art-image-group">
                            <img v-if="ArtImageUrl" :src="ArtImageUrl" alt="Uploaded Art Image" />
                            <div class="Art-placeholder-image" @click="addArtImage">
                                <div class="toUpload">
                                    <img src="../img/upload.png">
                                </div>
                                添加图片
                            </div>
                        </div>
                        <input type="file" ref="ArtImageUpload" accept="image/*" @change="handleArtImageChange"
                            style="display: none;" />
                    </div>

                    <div class="publish-button-row">
                        <button class="publish-button" @click="showNoticeModal = true">注册认证</button>
                    </div>
                    <label for="auto-create" class="publish-tip">数字艺术品将在审核通过后自动上架</label>
                </div>
            </div>
        </div>
        <!-- 用户须知弹窗 -->
        <div v-if="showNoticeModal" class="modal-publish">
            <div class="modal-content-publish">
                <span class="close-publish" @click="showNoticeModal = false">&times;</span>
                <h2>作品发布须知</h2>
                <p>欢迎使用 ArtChain 进行作品发布。在发布作品前，请您仔细阅读以下须知：</p>
                <ol>
                    <li>
                        <strong>原创确认</strong><br />
                        您保证所发布的作品是由您本人独立创作完成，不存在抄袭、剽窃等侵犯他人知识产权的行为。<br />
                        您确认对该作品拥有完整的版权，包括但不限于复制权、发行权、改编权等。
                    </li>
                    <li>
                        <strong>审核与存证</strong><br />
                        您提交的作品将经过平台的审核，审核内容包括作品的原创性、合法性以及信息的完整性。<br />
                        若作品审核成功，将自动进行区块链存证，确保作品信息的不可篡改和可追溯性。
                    </li>
                    <li>
                        <strong>作品上架与交易</strong><br />
                        审核成功并存证后的作品将自动上架，供其他用户浏览和购买。<br />
                        您所发布的作品在本平台以全权售出的方式进行交易，即买家将获得作品的完整版权。<br />
                        作品一经售出，具有唯一性，不可二次售出。
                    </li>
                    <li>
                        <strong>责任与义务</strong><br />
                        如因您发布的作品存在版权纠纷或其他法律问题，您将承担全部法律责任。<br />
                        请确保您提供的作品信息真实、准确、完整，否则可能导致作品审核不通过或交易失败。
                    </li>
                    <li>
                        <strong>平台权利</strong><br />
                        平台有权对您发布的作品进行审核、存证和管理，以确保平台的正常运行和用户的合法权益。<br />
                        平台有权根据法律法规和平台规则，对违规作品进行下架、删除等处理。
                    </li>
                </ol>
                <p>感谢您的理解与支持！</p>
                <div class="publish-but">
                    <button style="margin-right: 50px;" @click="showNoticeModal = false">取消</button>
                    <button @click="handlePublish">同意</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import axios from 'axios';

const router = useRouter();

// 用于存储藏品标题
const title = ref('');
// 用于存储藏品描述
const description = ref('');
// 用于存储藏品价格，新增价格的格式验证逻辑，这里简单示例只允许输入数字和小数点
const price = ref('');
const validatePrice = (inputPrice) => {
    // 使用正则表达式验证是否只包含数字和小数点，且小数点后最多两位（简单的价格格式验证示例）
    const priceRegex = /^\d+(\.\d{1,2})?$/;
    return priceRegex.test(inputPrice);
};

// 控制用户须知弹窗的显示状态
const showNoticeModal = ref(false);

const ArtImageUpload = ref(null);
const ArtImageUrl = ref('');
const ArtImageFile = ref(null); // 用于存储文件对象

const addArtImage = () => {
    ArtImageUpload.value.click();
};

const handleArtImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
        const reader = new FileReader();
        reader.onload = (event) => {
            ArtImageUrl.value = event.target.result; // 将图片 URL 存储到变量中
        };
        reader.readAsDataURL(file);
        ArtImageFile.value = file; // 存储文件对象
    }
};

const handlePublish = async () => {
    showNoticeModal.value = false;
    // 验证标题是否输入（去除两端空白字符后判断是否为空字符串）
    if (!title.value.trim()) {
        window.alert('请输入藏品标题！');
        return;
    }
    // 验证描述是否输入
    if (!description.value.trim()) {
        window.alert('请输入藏品描述！');
        return;
    }
    // 验证价格是否输入且格式是否正确
    if (!price.value.trim() || !validatePrice(price.value)) {
        window.alert('请输入正确的价格格式！');
        return;
    }
    // 验证是否上传了图片文件
    if (!ArtImageFile.value) {
        window.alert('请上传藏品图片！');
        return;
    }

    const token = localStorage.getItem('token'); // 从 Local Storage 获取 token

    const formData = new FormData();
    formData.append('title', title.value);
    formData.append('description', description.value);
    formData.append('tag', 'new'); // 如果有标签字段，可以在这里添加
    formData.append('artImage', ArtImageFile.value); // 上传的图片文件
    formData.append('price', price.value); // 将价格添加到表单数据中，发送到后端（假设后端接口接收该字段）

    try {
        const response = await axios.post('http://localhost:8001/add-art-piece', formData, {
            headers: {
                Authorization: token,
                'Content-Type': 'multipart/form-data' // 设置请求头
            }
        });

        if (response.data) {
            window.alert('艺术品上传成功！');
            setTimeout(() => {
                router.push('/Shop'); // 延时 1 秒后跳转
            }, 1000);
        }
    } catch (error) {
        if (error.response) {
            if (error.response.status === 401) {
                window.alert('未登录，请重新登录。'); // 提示用户未登录
                router.push('/Login'); // 重定向到登录页面
            } else {
                window.alert('艺术品上传失败，请重试。'); // 处理发布失败
            }
        } else {
            window.alert('网络错误，请稍后再试。'); // 处理网络错误
        }
    }
};

const publishStarsContainer = ref(null);
onMounted(() => {
    // 创建星星并初始化星星动画
    createPublishStars();
});
// 创建星尘
const createPublishStars = () => {
    const numberOfStars = 900;
    const styleSheet = document.styleSheets[0];

    // 预定义动画（所有星星共享一个动画）
    const keyframes = `
        @keyframes float {
            0% {
                transform: translateY(0px);
            }
            50% {
                transform: translateY(-20px);
            }
            100% {
                transform: translateY(0px);
            }
        }
    `;
    styleSheet.insertRule(keyframes, styleSheet.cssRules.length);

    for (let i = 0; i < numberOfStars; i++) {
        const star = document.createElement('div');
        star.className = 'publish-star';
        star.style.width = Math.random() * 1 + 1 + 'px';
        star.style.height = star.style.width;
        star.style.position = 'absolute';

        // 生成不同角度的路径
        const pathPositionX = Math.random() * 100;
        const pathPositionY1 = 43 + Math.sin(pathPositionX / 100 * Math.PI) * 20;
        const pathPositionY2 = 55 + Math.cos(pathPositionX / 200 * Math.PI) * 5;
        const pathPositionY3 = Math.random() * 10 + 55;
        const pathPositionY4 = Math.random() * 20 + 55;

        // 随机选择路径1、路径2或路径3
        const pathPositionY = Math.random() > 0.8 ? pathPositionY1 : (Math.random() > 0.8 ? pathPositionY2 : (Math.random() > 0.5 ? pathPositionY4 : pathPositionY3));

        // 设置星星位置
        star.style.left = pathPositionX + 'vw';
        star.style.top = pathPositionY + 'vh';

        // 动画持续时间和延迟
        star.style.animationDuration = Math.random() * 5 + 5 + 's';
        star.style.animationDelay = Math.random() * 3 + 3 + 's';

        // 随机透明度
        star.style.opacity = Math.random();

        // 为星星添加浮动动画（所有星星共享同一个动画）
        star.style.animation = `float ${Math.random() * 5 + 5}s ease-in-out infinite`;

        // 将星星添加到容器
        publishStarsContainer.value.appendChild(star);
    }
};
</script>

<style scoped>
@import url(../css/publish.css);
</style>