<script lang="ts">
  import { onMount } from 'svelte';
  import Footer from './lib/Footer.svelte';

  onMount(() => {
    console.log('app mounted');
    // replace "i" elements with data-feather with svg icons
    window.feather.replace();
  });
</script>

<main class="container">
  <Footer />
</main>

<style>

</style>