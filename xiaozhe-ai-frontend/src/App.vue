<template>
  <div class="app-container">
    <Sidebar />
    <div class="main-content">
      <ChatMain />
      <ChatInput />
    </div>
  </div>
</template>

<script setup>
import Sidebar from '@/components/layout/Sidebar.vue'
import ChatMain from '@/components/layout/ChatMain.vue'
import ChatInput from '@/components/layout/ChatInput.vue'
</script>

<style scoped>
.app-container {
  display: flex;
  height: 100vh;
  background: #fff;
}

.main-content {
  flex: 1;
  display: flex;
  flex-direction: column;
  background: #fff;
  position: relative;
}
</style>
