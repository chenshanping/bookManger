import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
      vue({
        refTransform: true,
        reactivityTransform: true,
      })
  ],
  server :{
    port: 8082,
      // proxy:{
      //   '/user':"http://11.0.1.60:8081"
      // }
  },
  base: './'
})
