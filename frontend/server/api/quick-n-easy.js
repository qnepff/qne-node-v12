import { defineEventHandler, readBody } from 'h3'
import fs from 'fs/promises'
import path from 'path'

const BASE_DIR = './data' // Base directory for file operations

export default defineEventHandler(async (event) => {
  const body = await readBody(event)
  const { action, fileName, content } = body

  try {
    switch (action) {
      case 'read':
        return await fs.readFile(path.join(BASE_DIR, fileName), 'utf-8')

      case 'write':
        await fs.writeFile(path.join(BASE_DIR, fileName), content)
        return { success: true, message: 'File written successfully' }

      case 'delete':
        await fs.unlink(path.join(BASE_DIR, fileName))
        return { success: true, message: 'File deleted successfully' }

      case 'list':
        const files = await fs.readdir(BASE_DIR)
        return files

      default:
        throw new Error('Invalid action')
    }
  } catch (error) {
    console.error('Error:', error)
    return { success: false, message: error instanceof Error ? error.message : String(error) }
  }
})