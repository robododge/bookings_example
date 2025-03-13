 (() => {
    function renderChart(chart, target) {
      const closestDivEl = target ?? document.currentScript.previousSibling;
  
      const chartData = closestDivEl.getAttribute("data-data");
      const chartLabels = closestDivEl.getAttribute("data-labels");
      const chartId = closestDivEl.getAttribute("data-id");
      const maxYvalue = closestDivEl.getAttribute("data-max-y-value");
      const maxY = maxYvalue ? parseInt(maxYvalue) : 1000;
      const myIndexAxis = closestDivEl.getAttribute("data-horizontal") ? 'y' : 'x';
      console.log("What was maxY?",maxY);
      console.log("What was chartData?",chartData);
      console.log("What was chartLabels?",chartLabels);

      if (chart) {
        console.log("Updating chart (1)!");
        chart.data.labels = JSON.parse(chartLabels);
        chart.data.datasets = [
          {
            ...chart.data.datasets[0],
            data: JSON.parse(chartData),
          },
        ];
        chart.update();
        console.log("Updated chart (2)!",chart.data);
        return;
      }

  
      const data = {
        labels: JSON.parse(chartLabels),
        datasets: [
          {
            data: JSON.parse(chartData),
            backgroundColor: (myIndexAxis === 'x')?["rgba(129, 57, 250, 0.4)"]:["rgba(218, 6, 53, 0.4)"],
            borderRadius: 4,
          },
        ],
      };
  

      chart = new Chart(document.getElementById(`chart-${chartId}`), {
        type: "bar",
        data: data,
        options: {
          plugins: {
            legend: {
              display: false,
            },
            tooltip: {
              enabled: true,
              backgroundColor: "#0a0a0a",
              titleColor: "white",
              bodyColor: "white",
              borderColor: "#262626",
              borderWidth: 1,
              padding: 10,
              cornerRadius: 4,
              displayColors: false,
              callbacks: {
                label: function (context) {
                  const value = context.raw;
                  return `${value} correctoins`;
                },
              },
            },
          },
          indexAxis: myIndexAxis,
          scales: {
            y: {
              beginAtZero: true,
              max: maxY,
              ticks: {
                display: true,
              },
              border: {
                display: false,
              },
              grid: {
                color: (myIndexAxis === 'y')? "transparent":"#262626",
              },
            },
            x: {
              ticks: {
                color: "#fff",
                format: (myIndexAxis === 'y')? {style: 'currency',
                        currency: 'USD'} : undefined,
              },
              grid: {
                color: (myIndexAxis === 'y')? "#262626": "transparent",
              },
            },
          },
        },
      });
  
      return {
        chartId,
        chart,
      };
    }
  
    const { chart, chartId } = renderChart();
  
    document.body.addEventListener("htmx:afterSettle", function (evt) {
      if (evt.detail.target.id === `chart-card-${chartId}`) {
        console.log("SHOULD be Rendering chart!-!-!-!");
        renderChart(chart, evt.detail.target);
      }
    });
  })();
