
<script>
// @ts-nocheck
    import { onMount } from 'svelte';
    import { MemberNode as MemberNodeClass} from '../lib/members/memberNode';
    import MemberNode from '../lib/members/MemberNode.svelte';

    let memberNode = null;

    // Fetch message from Go backend when the component is mounted
    async function fetchMessage(){
        const response = await fetch('http://localhost:8080/api/message');
        const resp = await response.json();
        console.log(resp);
        memberNode = new MemberNodeClass(resp)
        console.log(memberNode)
    }

    onMount( () => {
       fetchMessage();
    });
</script>

<div>
    
    <div>
        {#if memberNode}
            <MemberNode {memberNode}/>
        {:else}
            <p>loading...</p>
        {/if}
    </div>
</div>
