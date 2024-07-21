import { reactive, watch } from 'vue';

// 从localStorage中初始化状态
const user_mcu_id = localStorage.getItem('user_mcu_id') || '';

export const sharedState = reactive({
    user_mcu_id,
});

// 监视 user_mcu_id 的变化并将其存储在localStorage中
watch(
    () => sharedState.user_mcu_id,
    (newId) => {
        localStorage.setItem('user_mcu_id', newId);
    }
);
