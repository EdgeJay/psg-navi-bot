<script lang="ts">
  export let title = '';

  function handleBack() {
    history.back();
  }
</script>

<nav>
  <ul>
      <li>
          <a href="/" on:click|preventDefault={handleBack} class="secondary">
            <img src="/icons/arrow-left.svg" alt="Back" />
          </a>
      </li>
  </ul>
  <ul>
      <li><strong>{title}</strong></li>
  </ul>
</nav>