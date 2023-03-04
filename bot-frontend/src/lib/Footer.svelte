<script lang="ts">
  import { onMount } from 'svelte';
  import { appInfo } from '../stores/global';

  let appVersion = '';
  onMount(async () => {
    await appInfo.fetch();
    appVersion = $appInfo.ver;
  });
</script>

<footer class="version">
  Version {appVersion}
</footer>

<style>
  .version {
    font-size: 0.75rem;
    text-align: center;
    font-family: Arial, Helvetica, sans-serif;
    color: var(--tg-theme-hint-color, var(--fallback-hint-color));
  }
</style>