<script lang='ts'>
//@ts-nocheck
    import { onMount } from "svelte";
    import { createEventDispatcher } from "svelte";
    import type { MemberNode } from "./memberNode";

    export let memberNode: MemberNode;

    const dispatch = createEventDispatcher();

    // References to SVG elements
    let group: SVGGElement;
    let rect: SVGRectElement;
    let memberName: SVGTextElement;
    let memberDetails: SVGTextElement

    // Function to adjust rectangle size
    function adjustRectangleSize(){
        if (group && rect && memberName) {

            // Get bounding box of text elements
            const nameBBox = memberName.getBBox();
            const detailsBBox = memberDetails ? memberDetails.getBBox() : { width: 0, height: 0}

            // Calculate new dimensions
            const padding = Math.round(10);
            const width = Math.max(nameBBox.width, detailsBBox.width) + padding * 2;
            const height = nameBBox.height + detailsBBox.height + padding * 3;


            // const bbox = group.getBBox();
            // Update rect
            rect.setAttribute('width', Math.round(width).toString());
            rect.setAttribute('height', Math.round(height).toString());

            memberName.setAttribute('x', padding.toString());
            let memberY = Math.round(padding + nameBBox.height);
            memberName.setAttribute('y', memberY.toString());

            if (memberDetails) {
                memberDetails.setAttribute('x', padding.toString());
                let detailsY = Math.round(padding + nameBBox.height + detailsBBox.height);
                memberDetails.setAttribute('y', detailsY.toString());
            }
        }
    }

    function handleInteraction() {
        dispatch('nodeClick', {member:memberNode})
    }

    function handleKeydown() {
        if (event.key === 'Enter' || event.key === ' ') {
            event?.preventDefault();
            handleInteraction();
        }
    }

    // Call adjustRectangleSize when the component mounts
    onMount( () => {
        adjustRectangleSize();
    });
</script>

<!-- svelte-ignore a11y-no-static-element-interactions -->
<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<g 
    bind:this={group} 
    class="member-node" 
    on:click={handleInteraction}
    on:keydown={handleKeydown}
    tabindex="0"
    role="button"
    aria-label="View details for {memberNode.firstName} {memberNode.lastName}"
    >
    <rect bind:this={rect} class="node-background"/>

    <text bind:this={memberName} class="member-name">
        <tspan>{memberNode.firstName}</tspan>
        <tspan>{memberNode.lastName}</tspan>
    </text>
    <text bind:this={memberDetails} class="member-details">
        MiddleName: {memberNode.middleName}
    </text>
</g>

<style>
    .member-node {
        font-family: Arial, sans-serif;
        font-size: 14px;
        cursor: pointer;
    }

    .node-background{
        fill: #f0f0f0;
        stroke: #222;
        stroke-width: 2;
        rx: 10;
        ry: 10;
        opacity: 0.5;
    }
    .member-node:hover .node-background{
        fill: #df2a2a;
    }
    .member-name {
    font-weight: bold;
    }
    .member-details {
        font-size: 12px;
    }
</style>