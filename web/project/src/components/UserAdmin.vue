<template>
    <div class="user-admin-container">
        <!-- 顶部搜索区 -->
        <div class="action-bar">
            <div class="search-section">
                <el-input v-model="searchQuery" placeholder="输入用户ID、用户名" class="search-input" @keyup.enter="handleSearch">
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
                <span>用户 ID</span>
                <el-icon>
                    <ArrowDown />
                </el-icon>
            </div>
            <div class="header-item">
                <span>用户名</span>
                <el-icon>
                    <ArrowDown />
                </el-icon>
            </div>
            <div class="header-item">
                <span>密码</span>
                <el-icon>
                    <ArrowDown />
                </el-icon>
            </div>
            <div class="header-item">
                <span>私钥</span>
                <el-icon>
                    <ArrowDown />
                </el-icon>
            </div>
            <div class="header-item">
                <span>公钥</span>
                <el-icon>
                    <ArrowDown />
                </el-icon>
            </div>
            <div class="header-item">操作</div>
        </div>

        <!-- 表格内容 -->
        <div class="table-content">
            <div v-for="user in users" :key="user.uid" class="table-row">
                <div class="cell">
                    <el-icon>
                        <User />
                    </el-icon>
                    <span>{{ user.Uid }}</span>
                </div>
                <div class="cell">{{ user.Username }}</div>
                <div class="cell">{{ hidePassword(user.Userpass) }}</div>
                <div class="cell key-cell">{{ hideKey(user.PrivateKey) }}</div>
                <div class="cell key-cell">{{ hideKey(user.PublicKey) }}</div>
                <div class="cell actions">
                    <el-button-group>
                        <el-button class="action-btn edit" @click="handleEdit(user)">
                            <el-icon>
                                <Edit />
                            </el-icon>
                        </el-button>
                        <el-button class="action-btn delete" @click="handleDelete(user)">
                            <el-icon>
                                <Delete />
                            </el-icon>
                        </el-button>
                    </el-button-group>
                </div>
            </div>
        </div>

        <!-- 编辑用户弹窗 -->
        <el-dialog v-model="editDialogVisible" title="修改用户信息" class="custom-dialog" width="500px">
            <el-form :model="editingUser" label-width="100px">
                <el-form-item label="用户名">
                    <el-input v-model="editingUser.Username" placeholder="输入新用户名" />
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="editingUser.Userpass" placeholder="输入新密码" type="password" />
                </el-form-item>
            </el-form>
            <template #footer>
                <div class="dialog-footer">
                    <el-button type="primary" @click="saveEdit">保存</el-button>
                </div>
            </template>
        </el-dialog>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { Search, ArrowDown, User, Edit, Delete } from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from 'element-plus'

const searchQuery = ref('')
const editDialogVisible = ref(false)
const editingUser = ref({})
const users = ref([])

// 隐藏密码显示
const hidePassword = (password) => {
    return '*'.repeat(8)
}

// 隐藏密钥显示
const hideKey = (key) => {
    if (!key) return ''
    return key.substring(0, 6) + '...' + key.substring(key.length - 4)
}

// 获取所有用户
const fetchUsers = async () => {
    try {
        const response = await fetch('http://localhost:8001/admin/users', {
            headers: {
                'Authorization': localStorage.getItem('admin_token')
            }
        })
        if (response.ok) {
            users.value = await response.json()
        } else {
            window.alert('获取用户列表失败')
        }
    } catch (error) {
        window.alert('网络错误')
    }
}

// 搜索用户
const handleSearch = async () => {
    try {
        const response = await fetch(`http://localhost:8001/admin/users/search?q=${searchQuery.value}`, {
            headers: {
                'Authorization': localStorage.getItem('admin_token')
            }
        })
        if (response.ok) {
            users.value = await response.json()
        }
    } catch (error) {
        window.alert('搜索失败')
    }
}

// 处理编辑
const handleEdit = (user) => {
    editingUser.value = { ...user }
    editDialogVisible.value = true
}

// 保存编辑
const saveEdit = async () => {
    // 检查密码长度
    if (editingUser.value.Userpass.length < 6) {
        window.alert("密码长度不能小于6");
        return;
    }

    // 检查输入框非空
    if (!editingUser.value.Userpass || !editingUser.value.Username) {
        window.alert("输入不能为空");
        return;
    }

    try {
        const response = await fetch(`http://localhost:8001/admin/users/${editingUser.value.Uid}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': localStorage.getItem('admin_token')
            },
            body: JSON.stringify({
                username: editingUser.value.Username,
                password: editingUser.value.Userpass
            })
        })

        if (response.ok) {
            ElMessage.success('更新成功')
            editDialogVisible.value = false
            await fetchUsers()
        } else {
            const data = await response.json()
            window.alert(data.error || '更新失败')
        }
    } catch (error) {
        window.alert('网络错误')
    }
}

// 删除用户
const handleDelete = async (user) => {
    try {
        // 使用 window.confirm 替代 ElMessageBox
        const confirmed = window.confirm('确定要删除该用户吗？')

        if (!confirmed) {
            return // 如果点击取消，直接返回
        }

        const response = await fetch(`http://localhost:8001/admin/users/${user.Uid}`, {
            method: 'DELETE',
            headers: {
                'Authorization': localStorage.getItem('admin_token')
            }
        })

        if (response.ok) {
            ElMessage.success('删除成功')
            await fetchUsers()
        } else {
            ElMessage.error('删除失败')
        }
    } catch (error) {
        ElMessage.error('操作失败')
    }
}

// 页面加载时获取用户列表
onMounted(() => {
    fetchUsers()
})
</script>

<style scoped>
.user-admin-container {
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
    cursor: pointer;
    transition: all 0.3s;
}

.search-btn:hover {
    color: rgb(168, 206, 210, 0.9);
    background: rgba(168, 206, 210, 0.15);
    border-color: rgba(168, 206, 210, 0.3);
}

.table-header {
    display: grid;
    grid-template-columns: 0.8fr 1fr 1fr 1.5fr 1.55fr 90px;
    padding: 20px 24px;
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
    grid-template-columns: 0.8fr 1fr 1fr 1.5fr 1.5fr 100px;
    padding: 12px 24px;
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


.actions {
    justify-content: flex-end;
    margin-right: 25px;
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

.action-btn.edit {
    color: rgb(168, 206, 210, 0.7);
}

.action-btn.edit:hover {
    background-color: rgba(168, 206, 210, 0.05);
    border-color: rgba(168, 206, 210, 0.2);
}

.action-btn.delete {
    color: rgba(255, 107, 107, 0.7);
}

.action-btn.delete:hover {
    background-color: rgba(255, 107, 107, 0.05);
    border-color: rgba(255, 107, 107, 0.2);
}

/* 弹窗样式 */
:deep(.custom-dialog) {
    background-color: rgb(21, 26, 27);
    border-radius: 16px;
    border: 1px solid rgba(168, 206, 210, 0.428);
    backdrop-filter: blur(20px);
    margin-top: 25vh;
    margin-left: 33% !important;
}

:deep(.custom-dialog .el-dialog__header) {
    padding: 20px 24px;
    border-bottom: 1px solid rgba(168, 206, 210, 0.1);
}

:deep(.custom-dialog .el-dialog__title) {
    color: rgb(189, 208, 211);
    font-size: 18px;
    font-weight: 600;
}

:deep(.custom-dialog .el-dialog__body) {
    padding: 24px;
}

:deep(.custom-dialog .el-form-item__label) {
    color: rgba(168, 206, 210, 0.9);
    font-size: 16px;
    margin-top: 10px;
    margin-left: -42px;
    margin-right: 12px;
}

:deep(.custom-dialog .el-input__wrapper) {
    background: rgba(189, 208, 211, 0.1);
    box-shadow: none;
    padding: 0 12px; 
    margin-top: 10px;
    border: 1px solid rgba(168, 206, 210, 0.209);
    border-radius: 8px;
}

:deep(.custom-dialog .el-input__inner) {
    color: rgb(168, 206, 210, 0.8);
    font-size: 15px;
    height: 38px;
}

:deep(.custom-dialog .el-dialog__footer) {
    padding: 20px 24px;
    border-top: 1px solid rgba(168, 206, 210, 0.1);
}

:deep(.custom-dialog .el-button) {
    border-radius: 6px;
    padding: 5px 20px;
    font-weight: 500;
    margin-left: 30px;
}

:deep(.custom-dialog .el-button--primary) {
    height: 42px;
    background: rgba(168, 206, 210, 0.058);
    border: 1px solid rgba(168, 206, 210, 0.209);
    color: rgb(168, 206, 210, 0.85);
    padding: 0 20px;
    border-radius: 6px;
    font-size: 15px;
    font-weight: 600;
    transition: all 0.3s;
    cursor: pointer;
}

:deep(.custom-dialog .el-button--primary:hover) {
    color: rgb(168, 206, 210, 0.9);
    background: rgba(168, 206, 210, 0.15);
    border-color: rgba(168, 206, 210, 0.3);
}

/* 确认框样式 */
:deep(.el-message-box) {
    background-color: #1a1b23 !important;
    border: 1px solid rgba(168, 206, 210, 0.1);
    border-radius: 12px;
}

:deep(.el-message-box__header) {
    background-color: #1a1b23;
    padding: 20px 24px;
    border-bottom: 1px solid rgba(168, 206, 210, 0.1);
}

:deep(.el-message-box__title) {
    color: rgb(168, 206, 210) !important;
    font-size: 18px;
    font-weight: 500;
}

:deep(.el-message-box__content) {
    background-color: #1a1b23;
    color: rgb(168, 206, 210, 0.8) !important;
    padding: 24px;
}

:deep(.el-message-box__container) {
    background-color: #1a1b23;
}

:deep(.el-message-box__btns) {
    background-color: #1a1b23;
    padding: 12px 24px;
    border-top: 1px solid rgba(168, 206, 210, 0.1);
}

:deep(.el-message-box__btns .el-button) {
    background-color: rgba(40, 48, 49, 0.223);
    border: 1px solid rgba(168, 206, 210, 0.1);
    color: rgb(168, 206, 210);
    height: 36px;
    padding: 0 20px;
    font-size: 14px;
    font-weight: 500;
}

:deep(.el-message-box__btns .el-button--primary) {
    background-color: rgba(255, 107, 107, 0.1);
    border: 1px solid rgba(255, 107, 107, 0.2);
    color: rgb(255, 107, 107);
}

:deep(.el-message-box__btns .el-button:hover) {
    background-color: rgba(168, 206, 210, 0.1);
    border-color: rgba(168, 206, 210, 0.2);
}

:deep(.el-message-box__btns .el-button--primary:hover) {
    background-color: rgba(255, 107, 107, 0.15);
    border-color: rgba(255, 107, 107, 0.3);
}

:deep(.el-overlay) {
    background-color: rgba(0, 0, 0, 0.7);
}

/* 消息提示样式 */
:deep(.custom-message) {
    background-color: #1a1b23;
    border: 1px solid rgba(168, 206, 210, 0.1);
    color: rgb(168, 206, 210);
}
</style>