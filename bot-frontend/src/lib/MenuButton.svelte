<script lang="ts">
  export let label = '';
  export let iconType = '';
  export let icon = '';
</script>

<button>
  {#if iconType === 'devicons'}
    <i class="devicons devicons-{icon}"></i>
  {:else if iconType === 'feather'}
    <i data-feather="{icon}"></i>
  {/if}
  <span>{label}</span>
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

  button span {
    font-size: 1.25rem;
    cursor: pointer;
  }
</style>
