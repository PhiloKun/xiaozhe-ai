<template>
  <div class="sidebar">
    <!-- 顶部Logo区域 -->
    <div class="sidebar-header">
      <div class="logo-container">
        <div class="logo">小哲AI</div>
      </div>
      <div class="header-actions">
        <a-button type="text" class="action-btn">
          <template #icon>
            <IconFile/>
          </template>
        </a-button>
      </div>
    </div>

    <!-- 导航菜单 -->
    <div class="nav-menu">
      <div class="nav-item active">
        <div class="nav-icon">
          <IconPlus/>
        </div>
        <span class="nav-text">新建会话</span>
        <div class="nav-shortcut">Ctrl K</div>
      </div>

      <div class="nav-item">
        <div class="nav-icon">
          <IconRefresh/>
        </div>
        <span class="nav-text">小哲+</span>
      </div>

      <div class="nav-item">
        <div class="nav-icon ppt-icon">
          <IconFile/>
        </div>
        <span class="nav-text">PPT 助手</span>
      </div>

      <div class="nav-item">
        <div class="nav-icon">
          <IconClockCircle/>
        </div>
        <span class="nav-text">历史会话</span>
      </div>

      <div class="nav-item">
        <div class="nav-icon">
          <IconApps/>
        </div>
        <span class="nav-text">查看全部</span>
      </div>
    </div>

    <!-- 底部用户信息 -->
    <div class="user-section">
      <div class="user-info">
        <a-avatar class="user-avatar" :style="{ backgroundColor: '#f7ba2e' }">
          <IconUser/>
        </a-avatar>
        <span class="username">用户</span>
        <IconDown class="dropdown-icon"/>
      </div>
    </div>
  </div>
</template>

<script setup>
import {
  IconPlus,
  IconRefresh,
  IconFile,
  IconClockCircle,
  IconApps,
  IconUser,
  IconDown
} from '@arco-design/web-vue/es/icon'
</script>

<style scoped>
.sidebar {
  width: 280px;
  background: #2c2c2c;
  display: flex;
  flex-direction: column;
  height: 100vh;
  color: #fff;
}

.sidebar-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px 20px;
  border-bottom: 1px solid #404040;
}

.logo-container {
  display: flex;
  align-items: center;
}

.logo {
  color: #fff;
  width: 50px;
  height: 32px;
  border-radius: 4px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  font-size: 16px;
}

.header-actions {
  display: flex;
  align-items: center;
}

.action-btn {
  color: #fff;
  border: none;
  background: transparent;
}

.action-btn:hover {
  background: #404040;
}

.nav-menu {
  flex: 1;
  padding: 16px 0;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 12px 20px;
  margin: 4px 0;
  cursor: pointer;
  border-radius: 8px;
  margin: 4px 16px;
  position: relative;
}

.nav-item:hover {
  background: #404040;
}

.nav-item.active {
  background: #404040;
}

.nav-icon {
  width: 20px;
  height: 20px;
  margin-right: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
}

.ppt-icon {
  color: #ff4d4f;
}

.nav-text {
  flex: 1;
  font-size: 14px;
  color: #fff;
}

.nav-shortcut {
  background: #404040;
  color: #ccc;
  padding: 2px 8px;
  border-radius: 4px;
  font-size: 12px;
}

.user-section {
  padding: 16px 20px;
  border-top: 1px solid #404040;
}

.user-info {
  display: flex;
  align-items: center;
  cursor: pointer;
  padding: 8px;
  border-radius: 8px;
}

.user-info:hover {
  background: #404040;
}

.user-avatar {
  margin-right: 12px;
}

.username {
  flex: 1;
  font-size: 14px;
  color: #fff;
}

.dropdown-icon {
  color: #ccc;
  font-size: 12px;
}
</style>