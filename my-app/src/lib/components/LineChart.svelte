<script lang="ts">
    import { onMount } from 'svelte';
    import { Chart, LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend } from 'chart.js';

    // Register necessary Chart.js components
    Chart.register(LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend);

    let { statsArrOne, statsArrTwo, label }: { statsArrOne: number[], statsArrTwo: number[], label: string} = $props()

    let chart: Chart | null = null
    Chart.defaults.color = 'rgb(250,255,255)'
    Chart.defaults.font.size = 14
    let chartCanvas: HTMLCanvasElement


    onMount(() => {
        if (chart) { // update
            chart.destroy();
        }

        chart = new Chart(chartCanvas, {
            type: 'line',
            data: {
                labels: statsArrOne.map((_, i) => `Week ${i + 1}`),
                datasets: [
                    {
                        label: 'Alex ' + label,
                        data: statsArrOne,
                        borderColor: 'rgba(75, 192, 192, 1)',
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        fill: true,
                    },
                    {
                        label: 'Yoona ' + label,
                        data: statsArrTwo,
                        borderColor: 'rgba(255, 99, 132, 1)',
                        backgroundColor: 'rgba(255, 99, 132, 0.2)',
                        fill: true,
                    }
                ]
            },
            options: {
                responsive: true,
                scales: {
                    y: {
                        beginAtZero: true,
                        title: {
                            display: true,
                            text: label,
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: 'Weeks'
                        }
                    }
                },
            }
        });

        return () => {
            chart?.destroy();
        };
    });
</script>

<canvas class="w-full" bind:this={chartCanvas}></canvas>



