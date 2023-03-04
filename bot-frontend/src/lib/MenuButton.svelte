<script lang="ts">
  import { navigate } from 'svelte-routing';

  export let to = '#';
  export let label = '';
  export let iconType = '';
  export let icon = '';

  function handleClick() {
    navigate(to);
  }
</script>

<button on:click={handleClick}>
  {#if iconType === 'devicons'}
    <i class="devicons devicons-{icon}"></i>
  {:else if iconType === 'feather'}
    <i data-feather="{icon}"></i>
  {/if}
  <div class="label">
    <slot name="label">{label}</slot>
  </div>
</button>

<style>
  button {
    display: flex;
    flex-flow: row nowrap;
    align-items: center;
    justify-content: center;
  }

  button i {
    display: flex;
    flex-flow: column nowrap;
    justify-content: center;
    width: 3.5rem;
    height: 3.5rem;
    font-size: 3rem;
  }

  button .label {
    font-size: 1.25rem;
    cursor: pointer;
  }
</style>
