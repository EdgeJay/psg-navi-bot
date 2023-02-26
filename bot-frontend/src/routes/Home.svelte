<script lang="ts">
  import MenuButton from "../lib/MenuButton.svelte";
</script>

<section>
  <MenuButton to="/api/menu/dropbox" iconType="devicons" icon="dropbox">
    <span slot="label">Dropbox</span>
  </MenuButton>
  <MenuButton to="/api/menu/help" iconType="feather" icon="help-circle">
    <span slot="label">Ask me a question</span>
  </MenuButton>
</section>