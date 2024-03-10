import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import compression from 'vite-plugin-compression2';

// https://vitejs.dev/config/
export default defineConfig({
  base: '/static',
  plugins: [
    vue(),
    compression({
      threshold: 1024,
      deleteOriginalAssets: false,
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    }
  }
})
