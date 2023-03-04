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
    <img src="/icons/{icon}.svg" alt={label} />
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

  button img {
    margin: 0.4375rem 0;
    width: 2.5rem;
    height: 2.5rem;
    margin-right: 0.625rem;
    filter: invert(100%) sepia(1%) saturate(0%) hue-rotate(163deg) brightness(102%) contrast(102%);
  }

  button .label {
    font-size: 1.25rem;
    cursor: pointer;
  }
</style>
