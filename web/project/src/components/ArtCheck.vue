<template>
  <div class="art-check-container">
    <!-- 顶部搜索区 -->
    <div class="action-bar">
      <div class="search-section">
        <el-input v-model="searchQuery" placeholder="输入艺术品ID、标题、创作者、描述" class="search-input" @keyup.enter="handleSearch">
          <template #prefix>
            <el-icon>
              <Search />
            </el-icon>
          </template>
        </el-input>
        <el-button class="search-btn" @click="handleSearch">
          搜索
        </el-button>
      </div>
    </div>

    <!-- 表格头部 -->
    <div class="table-header">
      <div class="header-item">
        <span>艺术品 ID</span>
        <el-icon>
          <ArrowDown />
        </el-icon>
      </div>
      <div class="header-item">
        <span>标题</span>
        <el-icon>
          <ArrowDown />
        </el-icon>
      </div>
      <div class="header-item">
        <span>描述</span>
        <el-icon>
          <ArrowDown />
        </el-icon>
      </div>
      <div class="header-item">
        <span>创作者</span>
        <el-icon>
          <ArrowDown />
        </el-icon>
      </div>
      <div class="header-item">
        <span>标签</span>
        <el-icon>
          <ArrowDown />
        </el-icon>
      </div>
      <div class="header-item">
        <span>价格</span>
        <el-icon>
          <ArrowDown />
        </el-icon>
      </div>
      <div class="header-item">操作</div>
    </div>

    <!-- 表格内容 -->
    <div v-for="artwork in searchResults" :key="artwork.aid" class="table-row">
      <div class="cell">
        <el-icon>
          <Star />
        </el-icon>
        {{ artwork.aid }}
      </div>
      <div class="cell">{{ artwork.title }}</div>
      <div class="cell description">{{ artwork.description }}</div>
      <div class="cell">{{ artwork.creator_name }}</div>
      <div class="cell">{{ artwork.tag }}</div>
      <div class="cell price">¥{{ artwork.art_price.toFixed(2) }}</div>
      <div class="actions">
        <el-button class="action-btn view" @click="handleView(artwork)">
          <el-icon>
            <View />
          </el-icon>
        </el-button>
        <el-button class="action-btn approve" @click="handleApprove(artwork)">
          <el-icon>
            <Check />
          </el-icon>
        </el-button>
        <el-button class="action-btn reject" @click="handleReject(artwork)">
          <el-icon>
            <Close />
          </el-icon>
        </el-button>
      </div>
    </div>

    <!-- 详情弹窗 -->
    <el-dialog v-model="dialogVisible" title="艺术品详情" width="50%" class="custom-dialog" :before-close="handleClose">
      <div class="artwork-details" v-if="currentArtwork">
        <div class="artwork-image">
          <img :src="require(`../upload-img/${currentArtwork.art_image}`)" :alt="currentArtwork.title">
        </div>
        <div class="artwork-info">
          <div class="info-item">
            <div class="info-label">标题</div>
            <div class="info-content">{{ currentArtwork.title }}</div>
          </div>
          <div class="info-item">
            <div class="info-label">描述</div>
            <div class="info-content">{{ currentArtwork.description }}</div>
          </div>
          <div class="info-item">
            <div class="info-label">创作者</div>
            <div class="info-content">{{ currentArtwork.creator_name }}</div>
          </div>
          <div class="info-item">
            <div class="info-label">标签</div>
            <div class="info-content">{{ currentArtwork.tag }}</div>
          </div>
          <div class="info-item">
            <div class="info-label">价格</div>
            <div class="info-content">{{ currentArtwork.art_price }}</div>
          </div>
        </div>
      </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button class="dialog-btn cancel" @click="dialogVisible = false">取消</el-button>
          <el-button class="dialog-btn approve" @click="handleApprove(currentArtwork)">通过</el-button>
          <el-button class="dialog-btn reject" @click="handleReject(currentArtwork)">否决</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { Search, ArrowDown, View, Check, Close, Star } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'
import { ethers } from 'ethers'
import ArtCopyright from '@/contracts/ArtCopyright.json'

const router = useRouter()
const artworks = ref([])
const searchQuery = ref('')
const dialogVisible = ref(false)
const currentArtwork = ref(null)
const contract = ref(null)
const provider = ref(null)
const signer = ref(null)
const searchResults = ref([]) // 存储搜索结果

const handleSearch = () => {
  if (!searchQuery.value) {
    searchResults.value = artworks.value
    return
  }

  const query = searchQuery.value.toLowerCase()
  searchResults.value = artworks.value.filter(artwork =>
    artwork.aid.toString().includes(query) ||
    artwork.title.toLowerCase().includes(query) ||
    artwork.description.toLowerCase().includes(query) ||
    artwork.creator_name.toLowerCase().includes(query)
  )
}

// 初始化区块链连接
const initBlockchain = async () => {
  try {
    if (typeof window.ethereum === 'undefined') {
      throw new Error('请安装 MetaMask 钱包')
    }
    provider.value = new ethers.providers.Web3Provider(window.ethereum)
    signer.value = provider.value.getSigner()
    contract.value = new ethers.Contract(
      process.env.VUE_APP_CONTRACT_ADDRESS,
      ArtCopyright.abi,
      signer.value
    )
  } catch (error) {
    console.error('区块链初始化失败:', error)
    window.alert(error.message)
  }
}

// 获取待审核艺术品列表
const fetchPendingArtworks = async () => {
  try {
    const response = await fetch('http://localhost:8001/admin/pending-artworks', {
      headers: {
        'Authorization': localStorage.getItem('admin_token')
      }
    })
    if (response.ok) {
      artworks.value = await response.json()
      searchResults.value = artworks.value
    } else {
      window.alert('获取待审核艺术品失败')
    }
  } catch (error) {
    window.alert('网络错误')
  }
}

const handleView = (artwork) => {
  currentArtwork.value = artwork
  dialogVisible.value = true
}

const handleClose = () => {
  dialogVisible.value = false
  currentArtwork.value = null
}

const handleApprove = async (artwork) => {
  try {
    // 设置以太坊网络连接
    const provider = new ethers.providers.Web3Provider(window.ethereum);
    const signer = await provider.getSigner();
    // 智能合约实例化
    const contract = new ethers.Contract(
      process.env.VUE_APP_CONTRACT_ADDRESS,
      ArtCopyright.abi,
      signer
    );

    // 调用智能合约方法并发送交易
    const tx = await contract.registerArtwork(
      artwork.aid,
      artwork.title,                    // 标题
      artwork.description,              // 描述
      ethers.utils.parseEther(artwork.art_price.toString())  // 价格转换为 wei
    );

    // 等待交易确认并获取交易收据
    const receipt = await tx.wait();
    const contractHash = receipt.transactionHash;

    await handleReview(artwork.aid, 1, contractHash);
    window.alert('作品已上链并审核通过');
  } catch (error) {
    console.error('区块链操作失败:', error);
    console.error('Artwork data:', artwork); // 错误时打印数据
    window.alert('区块链操作失败，审核未完成');
  }
};

// 审核拒绝保持不变
const handleReject = async (artwork) => {
  if (window.confirm('确定拒绝该艺术品的审核吗？')) {
    await handleReview(artwork.aid, 0)
  }
}

// 修改处理审核的函数，添加 contractHash 参数
const handleReview = async (artId, status, contractHash = null) => {
  const adminToken = localStorage.getItem('admin_token')
  const adminId = localStorage.getItem('admin_id')

  if (!adminToken || !adminId) {
    window.alert('请先登录管理员账号')
    router.push('/AdminLog')
    return
  }

  try {
    const formData = new FormData()
    formData.append('art_id', artId)
    formData.append('status', status)
    formData.append('auditor_id', adminId)
    // 如果有合约标识，添加到表单数据中
    if (contractHash) {
      formData.append('contract_hash', contractHash)
    }

    const response = await fetch('http://localhost:8001/admin/review-artwork', {
      method: 'POST',
      headers: {
        'Authorization': adminToken
      },
      body: formData
    })

    if (response.ok) {
      window.alert('审核完成')
      dialogVisible.value = false
      await fetchPendingArtworks()
    } else {
      if (response.status === 401) {
        window.alert('登录已过期，请重新登录')
        localStorage.removeItem('admin_token')
        localStorage.removeItem('admin_id')
        router.push('/AdminLog')
        return
      }
      window.alert('审核操作失败')
    }
  } catch (error) {
    window.alert('网络错误')
  }
}

// 组件加载时初始化
onMounted(async () => {
  await initBlockchain()
  await fetchPendingArtworks()
})
</script>

<style scoped>
.art-check-container {
  padding: 24px;
  color: rgb(168, 206, 210);
  min-height: calc(100vh - 65px);
  background-color: #0f1314;
}

.action-bar {
  display: flex;
  justify-content: space-between;
  margin-bottom: 35px;
  padding: 0 5px;
}

.search-section {
  display: flex;
  align-items: center;
  gap: 12px;
}

.search-input {
  width: 420px;
}

.search-input :deep(.el-input__wrapper) {
  background-color: rgba(26, 34, 35, 0.5);
  box-shadow: none !important;
  border: 1px solid rgba(168, 206, 210, 0.25);
  border-radius: 6px;
  padding: 0 15px;
}

.search-input :deep(.el-input__wrapper:hover) {
  border-color: rgba(168, 206, 210, 0.3);
}

.search-input :deep(.el-input__wrapper.is-focus) {
  border-color: rgba(168, 206, 210, 0.4);
}

.search-input :deep(.el-input__inner) {
  color: rgb(168, 206, 210);
  height: 42px;
  font-size: 15px;
}

.search-input :deep(.el-input__prefix-inner) {
  color: rgb(168, 206, 210, 0.6);
}

.search-btn {
  height: 42px;
  background: rgba(168, 206, 210, 0.058);
  border: 1px solid rgba(168, 206, 210, 0.3);
  color: rgb(168, 206, 210, 0.85);
  padding: 0 20px;
  border-radius: 6px;
  font-size: 15px;
  font-weight: 600;
  transition: all 0.3s;
}

.search-btn:hover {
  color: rgb(168, 206, 210, 0.9);
  background: rgba(168, 206, 210, 0.15);
  border-color: rgba(168, 206, 210, 0.3);
}

.table-header {
  display: grid;
  grid-template-columns: 0.6fr 1.2fr 1.8fr 0.8fr 0.8fr 0.8fr 150px;
  padding: 16px 24px;
  background-color: rgba(40, 48, 49, 0.223);
  border-radius: 5px;
  margin-bottom: 3px;
}

.header-item {
  color: rgb(168, 206, 210, 0.8);
  font-size: 16px;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 8px;
}

.table-row {
  display: grid;
  grid-template-columns: 0.6fr 1.2fr 1.8fr 0.8fr 0.8fr 0.8fr 150px;
  padding: 16px 24px;
  background-color: rgba(40, 48, 49, 0.223);
  border-radius: 5px;
  margin-bottom: 3px;
  align-items: center;
  transition: all 0.3s;
}

.table-row:hover {
  background-color: rgba(26, 27, 35, 0.8);
  transform: translateX(5px);
}

.cell {
  color: rgb(189, 208, 211, 0.8);
  font-size: 15px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 10px;
}

.cell .el-icon {
  color: rgb(168, 206, 210);
  font-size: 18px;
}

/* 特别是描述列，确保它不会占用太多空间 */
.description {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  max-width: 70%;
}

.price {
  color: rgba(255, 246, 67, 0.6);
}

.actions {
  display: flex;
  gap: 2px;
  justify-content: flex-end;
  padding-right: 15px;
}

.action-btn {
  background-color: rgba(40, 48, 49, 0.223);
  border: 1px solid rgba(168, 206, 210, 0.1);
  width: 36px;
  height: 36px;
  padding: 0;
  border-radius: 8px;
  transition: all 0.3s;
}

.action-btn.view {
  color: rgb(168, 206, 210, 0.7);
}

.action-btn.view:hover {
  background-color: rgba(168, 206, 210, 0.05);
  border-color: rgba(168, 206, 210, 0.2);
  color: rgb(168, 206, 210);
}

.action-btn.approve {
  color: rgba(103, 194, 58, 0.7);
}

.action-btn.approve:hover {
  background-color: rgba(103, 194, 58, 0.05);
  border-color: rgba(103, 194, 58, 0.2);
  color: rgb(103, 194, 58);
}

.action-btn.reject {
  color: rgba(245, 108, 108, 0.7);
}

.action-btn.reject:hover {
  background-color: rgba(245, 108, 108, 0.05);
  border-color: rgba(245, 108, 108, 0.2);
  color: rgb(245, 108, 108);
}

/* 弹窗样式 */
:deep(.custom-dialog) {
  background-color: rgba(26, 34, 35, 0.627);
  border-radius: 12px;
  border: 1px solid rgba(168, 206, 210, 0.1);
  backdrop-filter: blur(20px);
  margin-left: 26% !important;
}

:deep(.custom-dialog .el-dialog__header) {
  padding: 20px 24px;
  border-bottom: 1px solid rgba(168, 206, 210, 0.1);
  margin-right: 0;
}

:deep(.custom-dialog .el-dialog__title) {
  color: rgb(189, 208, 211);
  font-size: 18px;
  font-weight: 600;
}

:deep(.custom-dialog .el-dialog__body) {
  padding: 24px;
}

.artwork-details {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.artwork-image {
  width: 100%;
  height: 400px;
  background-color: rgba(40, 48, 49, 0.5);
  border-radius: 8px;
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  border: 1px solid rgba(168, 206, 210, 0.1);
}

.artwork-image img {
  max-width: 100%;
  max-height: 100%;
  opacity: 0.9;
  object-fit: contain;
}

.artwork-info {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.info-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 16px;
  background-color: rgba(40, 48, 49, 0.3);
  border-radius: 8px;
  border: 1px solid rgba(168, 206, 210, 0.1);
}

.info-label {
  color: rgb(168, 206, 210, 0.7);
  font-size: 15px;
  font-weight: 500;
}

.info-content {
  color: rgb(168, 206, 210);
  font-size: 15px;
  line-height: 1.5;
}

:deep(.custom-dialog .el-dialog__footer) {
  padding: 20px 24px;
  border-top: 1px solid rgba(168, 206, 210, 0.1);
}

.dialog-footer {
  display: flex;
  gap: 12px;
  justify-content: flex-end;
}

.dialog-btn {
  height: 36px;
  padding: 0 20px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
  transition: all 0.3s;
}

.dialog-btn.cancel {
  background-color: rgba(40, 48, 49, 0.5);
  border: 1px solid rgba(168, 206, 210, 0.1);
  color: rgb(168, 206, 210);
}

.dialog-btn.cancel:hover {
  background-color: rgba(168, 206, 210, 0.1);
  border-color: rgba(168, 206, 210, 0.2);
}

.dialog-btn.approve {
  background-color: rgba(103, 194, 58, 0.1);
  border: 1px solid rgba(103, 194, 58, 0.2);
  color: rgb(103, 194, 58);
}

.dialog-btn.approve:hover {
  background-color: rgba(103, 194, 58, 0.15);
  border-color: rgba(103, 194, 58, 0.3);
}

.dialog-btn.reject {
  background-color: rgba(245, 108, 108, 0.1);
  border: 1px solid rgba(245, 108, 108, 0.2);
  color: rgb(245, 108, 108);
}

.dialog-btn.reject:hover {
  background-color: rgba(245, 108, 108, 0.15);
  border-color: rgba(245, 108, 108, 0.3);
}

/* 确认框样式 */
:deep(.el-message-box) {
  background-color: rgba(26, 34, 35, 0.95) !important;
  border: 1px solid rgba(168, 206, 210, 0.1);
  border-radius: 12px;
}

:deep(.el-message-box__header) {
  padding: 20px 24px;
  border-bottom: 1px solid rgba(168, 206, 210, 0.1);
}

:deep(.el-message-box__title) {
  color: rgb(168, 206, 210) !important;
  font-size: 18px;
  font-weight: 500;
}

:deep(.el-message-box__content) {
  padding: 24px;
  color: rgb(168, 206, 210, 0.8) !important;
  font-size: 15px;
}

:deep(.el-message-box__btns) {
  padding: 20px 24px;
  border-top: 1px solid rgba(168, 206, 210, 0.1);
}

:deep(.el-message-box__btns .el-button) {
  height: 36px;
  padding: 0 20px;
  border-radius: 8px;
  font-size: 14px;
  font-weight: 500;
}

/* 消息提示样式 */
:deep(.custom-message) {
  background-color: rgba(26, 34, 35, 0.95);
  border: 1px solid rgba(168, 206, 210, 0.1);
  color: rgb(168, 206, 210);
}
</style>