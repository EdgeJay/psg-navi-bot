<script lang="ts">
  import { onMount } from 'svelte';

  export let title = '';

  function handleBack() {
    history.back();
  }

  onMount(() => {
    // replace "i" elements with data-feather with svg icons
    window.feather.replace();
  });
</script>

<nav>
  <ul>
      <li>
          <a href="#" on:click={handleBack} class="secondary">
              <i data-feather="arrow-left"></i>
          </a>
      </li>
  </ul>
  <ul>
      <li><strong>{title}</strong></li>
  </ul>
</nav>