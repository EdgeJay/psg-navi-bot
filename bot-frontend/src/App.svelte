<script lang="ts">
  import { onMount } from 'svelte';
  import { Router, Route } from "svelte-routing";
  import Home from './routes/Home.svelte';
  import Dropbox from './routes/Dropbox.svelte';
  import DropboxAddFileRequest from './routes/DropboxAddFileRequest.svelte';
  import Footer from './lib/Footer.svelte';

  onMount(() => {
    console.log('app mounted');
    // replace "i" elements with data-feather with svg icons
    window.feather.replace();
  });
</script>

<main class="container">
  <Router>
    <Route path="dropbox"><Dropbox /></Route>
    <Route path="dbx-add-file-request"><DropboxAddFileRequest /></Route>
    <Route path="/"><Home /></Route>
  </Router>
  <Footer />
</main>

<style>

</style>