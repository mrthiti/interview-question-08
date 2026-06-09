<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { questionApi, type Question } from '@/services/api'

const router = useRouter()
const questions = ref<Question[]>([])
const loading = ref(false)
const error = ref('')

async function fetchQuestions() {
  try {
    loading.value = true
    const res = await questionApi.getAll()
    questions.value = res.data
  } catch {
    error.value = 'ไม่สามารถโหลดข้อมูลได้'
  } finally {
    loading.value = false
  }
}

async function deleteQuestion(id: number) {
  try {
    await questionApi.remove(id)
    await fetchQuestions()
  } catch {
    error.value = 'ไม่สามารถลบข้อมูลได้'
  }
}

onMounted(fetchQuestions)
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <div class="max-w-3xl mx-auto p-6">
      <div class="bg-green-600 text-white text-center py-3 rounded-t-lg font-bold text-lg">
        IT 08-1
      </div>

      <div class="bg-white shadow rounded-b-lg p-6">
        <button
          @click="router.push('/add')"
          class="mb-4 bg-green-500 hover:bg-green-600 text-white text-sm font-medium px-4 py-2 rounded transition-colors"
        >
          เพิ่มข้อสอบ
        </button>

        <p v-if="error" class="text-red-500 text-sm mb-3">{{ error }}</p>

        <p v-if="loading" class="text-gray-400 text-sm">กำลังโหลด...</p>

        <p v-else-if="!questions.length" class="text-gray-400 text-sm">ยังไม่มีข้อสอบ</p>

        <ul v-else class="space-y-2">
          <li
            v-for="q in questions"
            :key="q.ID"
            class="flex items-start gap-3 p-3 border border-gray-200 rounded"
          >
            <span class="mt-0.5 w-6 h-6 flex items-center justify-center rounded-full border-2 border-gray-400 text-xs text-gray-600 shrink-0">
              {{ q.order_number }}
            </span>
            <div class="flex-1">
              <p class="text-sm font-medium text-gray-800">{{ q.question }}</p>
              <ul class="mt-1 space-y-0.5">
                <li v-if="q.choice_1" class="flex items-center gap-2 text-sm text-gray-600">
                  <span class="w-4 h-4 rounded-full border border-gray-400 shrink-0"></span>{{ q.choice_1 }}
                </li>
                <li v-if="q.choice_2" class="flex items-center gap-2 text-sm text-gray-600">
                  <span class="w-4 h-4 rounded-full border border-gray-400 shrink-0"></span>{{ q.choice_2 }}
                </li>
                <li v-if="q.choice_3" class="flex items-center gap-2 text-sm text-gray-600">
                  <span class="w-4 h-4 rounded-full border border-gray-400 shrink-0"></span>{{ q.choice_3 }}
                </li>
                <li v-if="q.choice_4" class="flex items-center gap-2 text-sm text-gray-600">
                  <span class="w-4 h-4 rounded-full border border-gray-400 shrink-0"></span>{{ q.choice_4 }}
                </li>
              </ul>
            </div>
            <button
              @click="deleteQuestion(q.ID)"
              class="bg-red-500 hover:bg-red-600 text-white text-xs px-2 py-1 rounded transition-colors shrink-0"
            >
              ลบ
            </button>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>
