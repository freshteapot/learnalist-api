<script>
  import { replace } from "svelte-spa-router";
  import cache from "../lib/cache.js";
  import { putList, deleteList } from "../lib/api.js";
  import goto from "../lib/goto.js";
  import myLists from "../store/lists_by_me";
  import listsEdits from "../store/lists_edits.js";
  import Box from "./Box.svelte";
  import ListEditTitle from "./list_edit_title.svelte";
  import ListEditDataV1 from "./list_edit_data_v1.svelte";
  import ListEditDataV2 from "./list_edit_data_v2.svelte";
  import ListEditDataV3 from "./list_edit_data_v3.svelte";
  import ListEditDataV4 from "./list_edit_data_v4.svelte";
  import ListEditDataTodo from "./list_edit_data_todo.svelte";
  import ListEditLabels from "./list_edit_labels.svelte";

  export let aList;

  let listTypes = {
    v1: ListEditDataV1,
    v2: ListEditDataV2,
    v3: ListEditDataV3,
    v4: ListEditDataV4
  };

  let renderItem = Object.keys(listTypes)
    .filter(key => aList.info.type === key)
    .reduce((notFound, key) => {
      return listTypes[key];
    }, ListEditDataTodo);

  if (
    !aList.info.hasOwnProperty("interact") ||
    !aList.info.interact.hasOwnProperty("slideshow")
  ) {
    aList.info.interact = { slideshow: "0" };
  }

  $: canInteract = aList && aList.info.type === "v1";

  function cancel() {
    listsEdits.remove(aList.uuid);
    goto.list.view(aList.uuid);
  }

  async function remove() {
    const response = await deleteList(aList.uuid);
    if (response.status !== 200) {
      alert("failed try again");
      console.log("status from server was", response.status);
      return;
    }

    myLists.remove(aList.uuid);
    listsEdits.remove(aList.uuid);
    // TODO how to remove /lists/view/:uuid as well
    replace("/list/deleted");
  }

  async function save() {
    const response = await putList(aList);

    if (response.status !== 200) {
      alert("failed try again");
      console.log("status from server was", response.status);
      return;
    }

    try {
      listsEdits.remove(aList.uuid);
      myLists.update(aList);
      goto.list.view(aList.uuid);
    } catch (e) {
      alert("failed to clean up your edits");
    }
  }

  $: listsEdits.update(aList);
</script>

<Box>
  <ListEditTitle bind:title="{aList.info.title}" />
</Box>

<Box>
  <ListEditLabels bind:labels="{aList.info.labels}" />
</Box>

<Box>
  <svelte:component this="{renderItem}" bind:listData="{aList.data}" />
</Box>
<Box>
  <button on:click="{save}">Save</button>
  <button on:click="{cancel}">Cancel</button>
</Box>

<Box>
  <h1>Advanced</h1>
  <Box>
    <h2>Share</h2>
    <label>
      <input
        type="radio"
        bind:group="{aList.info.shared_with}"
        value="private"
      />
      Private
    </label>
    <label>
      <input
        type="radio"
        bind:group="{aList.info.shared_with}"
        value="public"
      />
      Public
    </label>
    <label>
      <input
        type="radio"
        bind:group="{aList.info.shared_with}"
        value="friends"
      />
      Friends
    </label>
  </Box>

  {#if canInteract}
    <Box>
      <h2>Interact</h2>
      <Box>
        <h3>Slideshow</h3>
        <label>
          <input
            type="radio"
            bind:group="{aList.info.interact.slideshow}"
            value="0"
          />
          Disable
        </label>

        <label>
          <input
            type="radio"
            bind:group="{aList.info.interact.slideshow}"
            value="1"
          />
          Enable
        </label>
      </Box>
    </Box>
  {/if}

  <Box>
    <h1>Danger</h1>
    <Box>
      <button on:click="{remove}">Delete this list forever</button>
    </Box>
  </Box>
</Box>