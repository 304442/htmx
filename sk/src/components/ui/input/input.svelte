<script lang="ts">
    import type { HTMLInputAttributes } from 'svelte/elements'
    import type { InputEvents } from '.'
  
    type $$Props = HTMLInputAttributes & { label: string }
    type $$Events = InputEvents
  
    let className: $$Props['class'] = undefined
    export let value: $$Props['value'] = undefined
    export let label = ''
    // export { className as class }
  </script>
  
  <div class="mb-4">
    <label
      for={$$restProps.name}
      class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
      >{label}</label
    >
    <input
      class="bg-gray-50 outline-none border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
      bind:value
      on:blur
      on:change
      on:click
      on:focus
      on:focusin
      on:focusout
      on:keydown
      on:keypress
      on:keyup
      on:mouseover
      on:mouseenter
      on:mouseleave
      on:paste
      on:input
      {...$$restProps}
    />
  </div>