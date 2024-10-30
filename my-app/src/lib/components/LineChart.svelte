<script lang="ts">
    import { onMount } from 'svelte';
    import { Chart, LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend } from 'chart.js';

    // Register necessary Chart.js components
    Chart.register(LineController, LineElement, PointElement, LinearScale, Title, CategoryScale, Tooltip, Legend);

    let { alexAccmScores, yoonaAccmScores }: { alexAccmScores: number[], yoonaAccmScores: number[]} = $props()

    let chart: Chart | null = null;
    let chartCanvas: HTMLCanvasElement;

    onMount(() => {
        if (chart) {
            chart.destroy();
        }

        chart = new Chart(chartCanvas, {
            type: 'line',
            data: {
                labels: alexAccmScores.map((_, i) => `Week ${i + 1}`),
                datasets: [
                    {
                        label: 'Alex Accumulated Scores',
                        data: alexAccmScores,
                        borderColor: 'rgba(75, 192, 192, 1)',
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',
                        fill: true,
                    },
                    {
                        label: 'Yoona Accumulated Scores',
                        data: yoonaAccmScores,
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
                            text: 'Scores'
                        }
                    },
                    x: {
                        title: {
                            display: true,
                            text: 'Weeks'
                        }
                    }
                }
            }
        });

        return () => {
            chart?.destroy();
        };
    });
</script>

<canvas class="w-full " bind:this={chartCanvas}></canvas>



