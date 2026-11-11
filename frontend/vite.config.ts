import react from '@vitejs/plugin-react'
import { defineConfig } from 'vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    tsconfigPaths: true,
  },
  server: {
    // strictPort: fail loudly on a collision rather than silently drifting to another
    // port (another project on this machine dev-serves on 5173).
    port: 5173,
    strictPort: true,
  },
})
