<script lang="ts">
    import { createEventDispatcher, onMount} from 'svelte';

    export let initialViewBox = { x: 0, y:0, width:1000, height:600};

    let containerDiv: HTMLDivElement;
    let svg: SVGSVGElement;
    let viewBox = {...initialViewBox};
    let isPanning = false;
    let startPoint = {x:0, y:0};
    let viewBoxStart = {x: 0, y:0};

    const dispatch = createEventDispatcher();

    const PAN_STEP = 40; // pan speed
    const ZOOM_FACTOR = 0.1 // zoom speed

    function startPan(event: MouseEvent){
        isPanning = true;
        startPoint = { x: event.clientX, y: event.clientY};
        viewBoxStart = { x: viewBox.x, y: viewBox.y};
    }

    function pan(event: MouseEvent){
        if (!isPanning) return;

        const dx = event.clientX - startPoint.x;
        const dy = event.clientY - startPoint.y;

        // Calculate the scaling factor based on the current zoom level
        const scaleFactor = viewBox.width / svg.clientWidth;

        viewBox.x = viewBoxStart.x - dx * scaleFactor;
        viewBox.y = viewBoxStart.y - dy * scaleFactor;

        updateViewBox();
    }

    function endPan() {
        isPanning = false;
    }

    function zoom(event: WheelEvent) {
        event.preventDefault();
        const zoomFactor = event.deltaY > 0 ? 1 + ZOOM_FACTOR : 1 - ZOOM_FACTOR;
        zoomToPoint(event.clientX, event.clientY, zoomFactor);

        const mouseX = event.clientX - svg.getBoundingClientRect().left;
        const mouseY = event.clientY - svg.getBoundingClientRect().top;

        const viewBoxX = viewBox.x + (mouseX / svg.clientWidth) * viewBox.width;
        const viewBoxY = viewBox.y + (mouseY / svg.clientHeight) * viewBox.height;
        

    }

    function zoomToPoint(clientX:number, clientY: number, zoomFactor: number) {
        const point = svg.createSVGPoint();
        point.x = clientX;
        point.y = clientY;
        const svgPoint = point.matrixTransform(svg.getScreenCTM()?.inverse());

        viewBox.width *= zoomFactor;
        viewBox.height *= zoomFactor;

        viewBox.x += (svgPoint.x - viewBox.x) * (1 - zoomFactor);
        viewBox.y += (svgPoint.y - viewBox.y) * (1 - zoomFactor);

        updateViewBox();
    }

    function updateViewBox() {
        svg.setAttribute('viewBox', `${viewBox.x} ${viewBox.y} ${viewBox.width} ${viewBox.height}`);
        dispatch('viewBoxUpdate', viewBox);
    }

    function handleKeydown(event: KeyboardEvent) {
        // Add keyboard controls for panning and zooming
        switch (event.key) {
            case 'ArrowUp':
                viewBox.y -= PAN_STEP;
                break;
            case 'ArrowDown':
                viewBox.y += PAN_STEP
                break;
                case 'ArrowLeft':
                viewBox.x += PAN_STEP;
                break;
            case 'ArrowRight':
                viewBox.x -= PAN_STEP;
                break;
            case '+':
            case '=':
                zoomToPoint(svg.clientWidth / 2, svg.clientHeight / 2, 1 - ZOOM_FACTOR);
                break;
            case '-':
            case '_':
                zoomToPoint(svg.clientWidth / 2, svg.clientHeight / 2, 1 + ZOOM_FACTOR);
                break;
            default:
                return; // Exit the function for other keys
        }
        event.preventDefault();
        updateViewBox();
    }

    onMount(() => {
        updateViewBox();
    });
</script>

<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
<!-- svelte-ignore a11y-no-noninteractive-element-interactions -->
<div
    bind:this={containerDiv}
    class="svg-container"
    tabindex="0"
    role="application"
    aria-label="interactive SVG with pan and zoom functionality"
    on:mousedown={startPan}
    on:mousemove={pan}
    on:mouseup={endPan}
    on:mouseleave={endPan}
    on:wheel={zoom}
    on:keydown={handleKeydown}
>
    <svg
        bind:this={svg}
        width="100%"
        height="100%"
    >
        <slot></slot>
    </svg>
</div>

<style>
    .svg-container {
        width: 100%;
        height: 100%;
        overflow: hidden;
        outline: none;
    }
    .svg-container:focus-visible{
        box-shadow: 0 0 0 2px #fff;
    }
</style>