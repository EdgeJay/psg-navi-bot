<script lang="ts">
  import { onMount } from 'svelte';
  import NavBar from "../lib/NavBar.svelte";
  import MenuButton from "../lib/MenuButton.svelte";

  onMount(() => {
    // replace "i" elements with data-feather with svg icons
    window.feather.replace();
  });
</script>

<section>
  <NavBar title="Dropbox" />
  <MenuButton to="dbx-add-file-request" iconType="feather" icon="plus-circle">
    <span slot="label">Add File Request</span>
  </MenuButton>
  <MenuButton to="dbx-list-file-requests" iconType="feather" icon="list">
    <span slot="label">List File Requests</span>
  </MenuButton>
</section>