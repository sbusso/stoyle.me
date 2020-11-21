<script>
    import { onMount } from 'svelte';

    export let newShort;

    let textarea;

    onMount(() => {
        textarea.select();
        document.execCommand('copy');
    });
</script>

<textarea bind:value={newShort} bind:this={textarea}></textarea>