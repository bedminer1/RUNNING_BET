<script lang="ts">
    import { onMount } from 'svelte';
    import { Chart, LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend } from 'chart.js';

    // Register necessary Chart.js components
    Chart.register(LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend);

    let { stats, label }: { stats: DataSet[], label: string} = $props()

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
                labels: stats[0].data.map((_, i) => `Week ${i + 1}`),
                datasets: stats.map(stat => ({
                    label: stat.label + " " + label,
                    data: stat.data,
                    borderColor: stat.borderColor || 'rgba(75, 192, 192, 1)',
                    backgroundColor: stat.backgroundColor || 'rgba(75, 192, 192, 0.2)',
                    fill: true,
                }))
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



