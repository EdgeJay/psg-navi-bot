<script lang="ts">
  import NavBar from "../lib/NavBar.svelte";
  import MenuButton from "../lib/MenuButton.svelte";
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