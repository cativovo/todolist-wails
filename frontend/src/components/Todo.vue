<script setup lang="ts">
import Button from "./shadcn/button/Button.vue";
import Checkbox from "./shadcn/checkbox/Checkbox.vue";
type Props = {
  // Go type: time
  created_at: any;
  // Go type: time
  updated_at: any;
  id: string;
  title: string;
  completed: boolean;
};

type Emits = {
  delete: [id: string];
  "toggle-check": [id: string, checked: boolean];
  "update-title": [id: string, title: string];
};

const props = defineProps<Props>();
const emit = defineEmits<Emits>();
</script>

<template>
  <div class="flex items-center space-x-2">
    <Checkbox
      :checked="props.completed"
      @update:checked="(checked) => emit('toggle-check', props.id, checked)"
    />
    <div
      class="inline-block overflow-hidden p-1 w-96 h-9 whitespace-nowrap text-ellipsis"
    >
      <input
        v-if="!props.completed"
        type="text"
        class="p-1 w-full"
        :value="props.title"
        @input="
          emit(
            'update-title',
            props.id,
            ($event.target as HTMLInputElement).value,
          )
        "
      />
      <span v-else class="p-1 line-through">{{ props.title }}</span>
    </div>
    <Button variant="ghost" @click="emit('delete', props.id)">x</Button>
  </div>
</template>
