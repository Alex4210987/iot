import { createApp } from 'vue'
import ECharts from 'vue-echarts'
import {use} from 'echarts/core'
import router from './router/router'
import App from './App.vue'

// 手动引入 ECharts 各模块来减小打包体积
import {
    CanvasRenderer
} from 'echarts/renderers'
import {
    BarChart
} from 'echarts/charts'
import {
    GridComponent,
    TooltipComponent
} from 'echarts/components'

use([
    CanvasRenderer,
    BarChart,
    GridComponent,
    TooltipComponent
]);

const app = createApp(App).use(router).mount('#app')

