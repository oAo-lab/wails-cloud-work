import {defineConfig} from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    // 服务器地址
    port: 7000,
    // 代理地址
    // proxy: {
    //   '/api/v1': {
    //     target: 'http://localhost:50010',
    //     ws: false,
    //     changeOrigin: true,
    //     rewrite: (path) => path.replace(/^\/api\/v1/, '')
    //   },
    // },
  },
})
