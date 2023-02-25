<script lang="ts">
  import { appVersion } from '../stores/global';
</script>

<footer class="version">
  Version {$appVersion}
</footer>

<style>
  .version {
    font-size: 0.75rem;
    text-align: center;
    font-family: Arial, Helvetica, sans-serif;
    color: var(--tg-theme-hint-color, var(--fallback-hint-color));
  }
</style>