<script lang="ts">
  import { onMount } from 'svelte';
  import { Router, Link, Route } from "svelte-routing";
  import Home from './routes/Home.svelte';
  import Footer from './lib/Footer.svelte';

  onMount(() => {
    console.log('app mounted');
    // replace "i" elements with data-feather with svg icons
    window.feather.replace();
  });
</script>

<main class="container">
  <Router>
    <Route path="/"><Home /></Route>
  </Router>
  <Footer />
</main>

<style>

</style>