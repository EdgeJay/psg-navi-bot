<script lang="ts">
  export let title = '';

  let colorScheme = window.Telegram.WebApp.colorScheme;

  function handleBack() {
    history.back();
  }
</script>

<nav>
  <ul>
      <li>
          <a href="/" on:click|preventDefault={handleBack} class="secondary">
            <img class={colorScheme === 'dark' ? 'filter-to-white' : ''} src="/icons/arrow-left.svg" alt="Back" />
          </a>
      </li>
  </ul>
  <ul>
      <li><strong>{title}</strong></li>
  </ul>
</nav>

<style>
  .filter-to-white {
    filter: invert(100%) sepia(1%) saturate(0%) hue-rotate(163deg) brightness(102%) contrast(102%);
  }
</style>
