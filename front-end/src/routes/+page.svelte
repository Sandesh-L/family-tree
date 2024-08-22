
<script>
    import { onMount } from 'svelte';

    let message = "";

    // Fetch message from Go backend when the component is mounted
    async function fetchMessage(){
        const response = await fetch('http://localhost:8080/api/message');
        const data = await response.json();
        message = data.message;
    }

    onMount( () => {
        fetchMessage();
    });
</script>

<div>
    <h1>Welcome to SvelteKit</h1>
    
    <div>
        <h1> Message from Go Backend: </h1> 
        <p>{message}</p>
    </div>
</div>

<p>Visit <a href="https://kit.svelte.dev">kit.svelte.dev</a> to read the documentation</p>
