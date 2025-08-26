<template>
    <div class="chat-main">
        <!-- 标题区域 -->
        <div class="title-section">
            <h1 class="main-title">小哲</h1>
        </div>

        <!-- 输入区域 -->
        <div class="input-section">
            <div class="input-container">
                <a-textarea v-model="inputText" placeholder="尽管问..." :auto-size="{ minRows: 1, maxRows: 6 }"
                    class="chat-input" @keydown.ctrl.enter="handleSend" />
                <div class="input-actions">
                    <div class="action-buttons">
                                                 <a-button type="text" class="action-btn">
                             <template #icon>
                                 <icon-refresh />
                             </template>
                         </a-button>
                        <a-button type="text" class="action-btn">
                            <template #icon>
                                <icon-search />
                            </template>
                            <span>深度研究</span>
                        </a-button>
                        <div class="model-selector">
                            <span>小哲1.5</span>
                            <icon-down />
                        </div>
                                                 <a-button type="text" class="action-btn">
                             <template #icon>
                                 <icon-attachment />
                             </template>
                         </a-button>
                         <a-button type="text" class="action-btn">
                             <template #icon>
                                 <icon-attachment />
                             </template>
                         </a-button>
                        <a-button type="primary" class="send-btn" @click="handleSend">
                            <template #icon>
                                <icon-send />
                            </template>
                        </a-button>
                    </div>
                </div>
            </div>
        </div>

        <!-- 建议按钮区域 -->
        <div class="suggestions-section">
            <div class="suggestions-grid">
                <div class="suggestion-item">
                    <div class="suggestion-icon colorful">
                        <icon-star />
                    </div>
                    <span>小哲 K2 发布并开源,模型即 Agent</span>
                </div>
                <div class="suggestion-item">
                    <div class="suggestion-icon">
                        <icon-star />
                    </div>
                    <span>申请加入 小哲</span>
                </div>
                                 <div class="suggestion-item">
                     <div class="suggestion-icon">
                         <icon-refresh />
                     </div>
                     <span>轻松创作无限流小说</span>
                 </div>
                                 <div class="suggestion-item">
                     <div class="suggestion-icon">
                         <icon-star />
                     </div>
                     <span>最新研究:AI不会讲笑话</span>
                 </div>
                                 <div class="suggestion-item">
                     <div class="suggestion-icon">
                         <icon-star />
                     </div>
                     <span>历代王朝开国难度排行榜</span>
                 </div>
            </div>
        </div>

        <!-- 底部提示 -->
        <div class="footer-note">
            <span>内容由AI生成,请仔细甄别</span>
        </div>
    </div>
</template>

<script setup>
import { ref } from 'vue'
import { 
  IconRefresh, 
  IconSearch, 
  IconDown, 
  IconAttachment, 
  IconSend,
  IconStar
} from '@arco-design/web-vue/es/icon'

const inputText = ref('')

const handleSend = () => {
  if (!inputText.value.trim()) return
  // 处理发送逻辑
  console.log('发送消息:', inputText.value)
  inputText.value = ''
}
</script>

<style scoped>
.chat-main {
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    padding: 40px 20px;
    max-width: 800px;
    margin: 0 auto;
    width: 100%;
}

.title-section {
    margin-bottom: 40px;
    text-align: center;
}

.main-title {
    font-size: 48px;
    font-weight: bold;
    color: #000;
    margin: 0;
}

.input-section {
    width: 100%;
    margin-bottom: 40px;
}

.input-container {
    background: #fff;
    border: 1px solid #e5e5e5;
    border-radius: 12px;
    padding: 16px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.chat-input {
    border: none;
    resize: none;
    font-size: 16px;
    line-height: 1.5;
    margin-bottom: 12px;
}

.chat-input:focus {
    outline: none;
    box-shadow: none;
}

.input-actions {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.action-buttons {
    display: flex;
    align-items: center;
    gap: 8px;
}

.action-btn {
    color: #666;
    border: none;
    background: transparent;
    padding: 4px 8px;
    border-radius: 4px;
}

.action-btn:hover {
    background: #f5f5f5;
}

.model-selector {
    display: flex;
    align-items: center;
    gap: 4px;
    color: #666;
    font-size: 14px;
    cursor: pointer;
    padding: 4px 8px;
    border-radius: 4px;
}

.model-selector:hover {
    background: #f5f5f5;
}

.send-btn {
    background: #165dff;
    border: none;
    border-radius: 50%;
    width: 32px;
    height: 32px;
    display: flex;
    align-items: center;
    justify-content: center;
}

.send-btn:hover {
    background: #0e4bcc;
}

.suggestions-section {
    width: 100%;
    margin-bottom: 40px;
}

.suggestions-grid {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
    gap: 12px;
}

.suggestion-item {
    display: flex;
    align-items: center;
    gap: 12px;
    padding: 12px 16px;
    background: #fff;
    border: 1px solid #e5e5e5;
    border-radius: 8px;
    cursor: pointer;
    transition: all 0.2s;
}

.suggestion-item:hover {
    border-color: #165dff;
    background: #f8f9ff;
}

.suggestion-icon {
    width: 24px;
    height: 24px;
    display: flex;
    align-items: center;
    justify-content: center;
    color: #666;
}

.suggestion-icon.colorful {
    color: #165dff;
}

.suggestion-item span {
    font-size: 14px;
    color: #333;
    flex: 1;
}

.footer-note {
    text-align: center;
    color: #999;
    font-size: 12px;
    margin-top: auto;
}
</style>