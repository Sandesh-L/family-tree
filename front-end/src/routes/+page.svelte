
<script lang='ts'>
    import {onMount} from 'svelte';
    import { MemberNode as MemberNodeClass} from '../lib/members/memberNode';
    import MemberNode from '../lib/members/MemberNode.svelte';
    import InteractiveSvgContainer from '$lib/InteractiveSVGContainer.svelte';

    let memberNode: MemberNodeClass | null = null;
    let currentViewBox = {x:0, y:0, width:500, height:500};
    let selectedMember: MemberNodeClass | null = null;

    // Fetch message from Go backend when the component is mounted
    async function fetchMessage(){
        const response = await fetch('http://localhost:8080/api/message');
        const resp = await response.json();
        console.log(resp);
        memberNode = new MemberNodeClass(resp)
        console.log(memberNode)
    }

    function handleViewBoxUpdate(event: CustomEvent) {
        currentViewBox = event.detail;
    }

    function handleNodeClick(event: CustomEvent) {
        selectedMember = event.detail.member;
        console.log('Selected member: ', selectedMember);
    }

    onMount( () => {
       fetchMessage();
    });
</script>

<InteractiveSvgContainer initialViewBox={currentViewBox} on:viewBoxUpdate={handleViewBoxUpdate}>
    {#if memberNode}
        <MemberNode {memberNode} on:nodeClick={handleNodeClick}/>
    {:else}
        <text x='10' y='20' font-family='Arial, sans-serif' font-size='16'> Loading... </text>
    {/if}
</InteractiveSvgContainer>

<p>Current ViewBox: x={currentViewBox.x}, y={currentViewBox.y}, width={currentViewBox.width}, height={currentViewBox.height}</p>
