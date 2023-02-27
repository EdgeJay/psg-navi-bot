<script lang="ts">
  import { appInfo } from '../stores/global';
</script>

<footer class="version">
  Version {$appInfo ? $appInfo.ver : ''}
</footer>

<style>
  .version {
    font-size: 0.75rem;
    text-align: center;
    font-family: Arial, Helvetica, sans-serif;
    color: var(--tg-theme-hint-color, var(--fallback-hint-color));
  }
</style>