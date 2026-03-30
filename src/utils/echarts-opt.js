import * as echarts from 'echarts'

const COLORS = {
  primary: '#14b8a6',   // teal-500
  secondary: '#3b82f6', // blue-500
  warning: '#f59e0b',   // amber-500
  success: '#10b981',   // emerald-500
  danger: '#ef4444',    // red-500
  textMain: '#1f2937',  // gray-800
  textSub: '#6b7280',   // gray-500
  gridLine: '#e5e7eb',  // gray-200
  bgHover: '#f3f4f6'    // gray-100
};

/**
 * 获取雷达图配置
 * @param {Array} indicators - 指示器数组 [{ name, max }]
 * @param {Array} dataValues - 数据值数组 [80, 90, ...]
 * @param {String} seriesName - 系列名称
 */
export const getRadarOption = (indicators, dataValues, seriesName = '能力评估') => {
  return {
    tooltip: {
      trigger: 'item',
      backgroundColor: 'rgba(255, 255, 255, 0.9)',
      borderColor: COLORS.gridLine,
      borderWidth: 1,
      textStyle: { color: COLORS.textMain },
      extraCssText: 'box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px;'
    },
    legend: {
      data: [seriesName],
      bottom: '0%',
      textStyle: { color: COLORS.textSub }
    },
    radar: {
      indicator: indicators,
      shape: 'circle', // 圆形雷达图更现代
      center: ['50%', '50%'],
      radius: '65%',
      axisName: {
        color: COLORS.textSub,
        fontWeight: 'bold',
        fontSize: 12,
        padding: [3, 5]
      },
      splitLine: {
        lineStyle: {
          color: COLORS.gridLine,
          width: 1
        }
      },
      splitArea: {
        show: true,
        areaStyle: {
          color: ['#ffffff', '#f9fafb'] // 交替背景色
        }
      },
      axisLine: {
        lineStyle: {
          color: COLORS.gridLine
        }
      }
    },
    series: [
      {
        name: seriesName,
        type: 'radar',
        data: [
          {
            value: dataValues,
            name: seriesName,
            symbol: 'circle',
            symbolSize: 8,
            itemStyle: {
              color: COLORS.primary,
              shadowBlur: 10,
              shadowColor: 'rgba(20, 184, 166, 0.5)'
            },
            lineStyle: {
              width: 3,
              color: COLORS.primary
            },
            areaStyle: {
              color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                { offset: 0, color: 'rgba(20, 184, 166, 0.5)' },
                { offset: 1, color: 'rgba(20, 184, 166, 0.1)' }
              ])
            }
          }
        ]
      }
    ]
  };
};

/**
 * 获取对比雷达图配置 (个人 vs 岗位)
 * @param {Array} indicators
 * @param {Array} personalData
 * @param {Array} jobRequirementData
 */
export const getComparisonRadarOption = (indicators, personalData, jobRequirementData) => {
  return {
    tooltip: {
      trigger: 'item',
      backgroundColor: 'rgba(255, 255, 255, 0.92)',
      borderColor: COLORS.gridLine,
      borderWidth: 1,
      textStyle: { color: COLORS.textMain },
      extraCssText: 'box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08); border-radius: 10px;'
    },
    legend: {
      data: ['我的能力', '岗位要求'],
      bottom: '0%',
      textStyle: { color: COLORS.textSub }
    },
    radar: {
      indicator: indicators,
      shape: 'circle',
      center: ['50%', '50%'],
      radius: '65%',
      axisName: {
        color: COLORS.textSub,
        fontWeight: 'bold'
      },
      splitLine: { lineStyle: { color: COLORS.gridLine } },
      splitArea: { show: false },
      axisLine: { lineStyle: { color: COLORS.gridLine } }
    },
    series: [
      {
        name: '能力对比',
        type: 'radar',
        data: [
          {
            value: personalData,
            name: '我的能力',
            itemStyle: { color: COLORS.primary },
            areaStyle: { color: 'rgba(20, 184, 166, 0.3)' },
            lineStyle: { width: 2 }
          },
          {
            value: jobRequirementData,
            name: '岗位要求',
            itemStyle: { color: COLORS.warning },
            lineStyle: { type: 'dashed', width: 2, color: COLORS.warning },
            areaStyle: { opacity: 0 } // 不填充背景，只显虚线
          }
        ]
      }
    ]
  };
};

// 如果需要柱状图、折线图等，可以继续在此文件中添加函数
export const getBarOption = (categories, data, title) => {
  return {
    tooltip: { trigger: 'axis' },
    grid: { left: '3%', right: '4%', bottom: '3%', containLabel: true },
    xAxis: {
      type: 'category',
      data: categories,
      axisLine: { lineStyle: { color: COLORS.gridLine } },
      axisLabel: { color: COLORS.textSub }
    },
    yAxis: {
      type: 'value',
      splitLine: { lineStyle: { color: COLORS.gridLine, type: 'dashed' } },
      axisLabel: { color: COLORS.textSub }
    },
    series: [{
      data: data,
      type: 'bar',
      barWidth: '40%',
      itemStyle: {
        color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
          { offset: 0, color: COLORS.primary },
          { offset: 1, color: '#99f6e4' } // teal-200
        ]),
        borderRadius: [4, 4, 0, 0]
      }
    }]
  };
};

export const getJobTransitionGraphOption = ({ nodes, links, title }) => {
  return {
    title: title
      ? {
        text: title,
        left: 'center',
        top: 6,
        textStyle: { color: COLORS.textMain, fontSize: 14, fontWeight: 'bold' },
      }
      : undefined,
    tooltip: {
      trigger: 'item',
      backgroundColor: 'rgba(255, 255, 255, 0.92)',
      borderColor: COLORS.gridLine,
      borderWidth: 1,
      textStyle: { color: COLORS.textMain },
      extraCssText: 'box-shadow: 0 8px 24px rgba(15, 23, 42, 0.08); border-radius: 10px;',
    },
    series: [
      {
        type: 'graph',
        layout: 'force',
        roam: true,
        draggable: true,
        force: {
          repulsion: 220,
          edgeLength: [90, 160],
          gravity: 0.08,
        },
        data: nodes.map((n) => ({
          ...n,
          symbolSize: n.symbolSize ?? (n.isTarget ? 58 : n.isCurrent ? 52 : 44),
          itemStyle: {
            color: n.isTarget ? COLORS.primary : n.isCurrent ? '#2dd4bf' : '#99f6e4',
            borderColor: '#ffffff',
            borderWidth: 2,
            shadowBlur: 12,
            shadowColor: 'rgba(20, 184, 166, 0.18)',
          },
          label: {
            show: true,
            color: COLORS.textMain,
            fontWeight: 'bold',
          },
        })),
        links: links.map((l) => ({
          ...l,
          lineStyle: {
            color: 'rgba(20, 184, 166, 0.55)',
            width: Math.max(1, Math.min(6, Number(l.value ?? 2))),
            curveness: 0.22,
          },
        })),
        edgeSymbol: ['none', 'arrow'],
        edgeSymbolSize: 10,
        label: { show: false },
      },
    ],
  }
}



export const getHeatGaugeOption = ({ value, title, subtitle }) => {
  const v = Math.max(0, Math.min(100, Number(value) || 0))
  return {
    title: {
      show: Boolean(title || subtitle),
      text: title || '',
      subtext: subtitle || '',
      left: 'center',
      top: 6,
      textStyle: { color: COLORS.textMain, fontSize: 14, fontWeight: 'bold' },
      subtextStyle: { color: COLORS.textSub, fontSize: 11 },
    },
    series: [
      {
        type: 'gauge',
        center: ['50%', '58%'],
        startAngle: 210,
        endAngle: -30,
        min: 0,
        max: 100,
        splitNumber: 5,
        radius: '92%',
        progress: {
          show: true,
          width: 14,
          roundCap: true,
          itemStyle: {
            color: new echarts.graphic.LinearGradient(0, 0, 1, 0, [
              { offset: 0, color: '#99f6e4' },
              { offset: 1, color: COLORS.primary },
            ]),
          },
        },
        axisLine: {
          lineStyle: {
            width: 14,
            color: [[1, 'rgba(148, 163, 184, 0.25)']],
          },
        },
        axisTick: { show: false },
        splitLine: { show: false },
        axisLabel: { show: false },
        pointer: { show: false },
        anchor: { show: false },
        detail: {
          valueAnimation: true,
          offsetCenter: [0, 8],
          fontSize: 34,
          fontWeight: 'bold',
          color: COLORS.textMain,
          formatter: '{value}',
        },
        title: {
          offsetCenter: [0, 38],
          fontSize: 12,
          color: COLORS.textSub,
        },
        data: [{ value: v, name: '热度指数' }],
      },
    ],
  }
}

export const getRoiBubbleOption = ({ points, title }) => {
  const safe = Array.isArray(points) ? points : []
  const data = safe.map((p) => {
    const effort = Math.max(0, Math.min(10, Number(p.effort) || 0))
    const impact = Math.max(0, Math.min(10, Number(p.impact) || 0))
    const roi = Math.max(0, Math.min(10, Number(p.roi) || 0))
    return {
      name: p.name,
      value: [effort, impact, roi],
      effort,
      impact,
      roi,
      raw: p,
    }
  })

  return {
    title: title
      ? {
        text: title,
        left: 0,
        top: 0,
        textStyle: { color: COLORS.textMain, fontSize: 14, fontWeight: 'bold' },
      }
      : undefined,
    grid: { left: 28, right: 18, top: title ? 40 : 20, bottom: 26 },
    tooltip: {
      trigger: 'item',
      backgroundColor: 'rgba(255, 255, 255, 0.92)',
      borderColor: COLORS.gridLine,
      borderWidth: 1,
      textStyle: { color: COLORS.textMain },
      extraCssText: 'box-shadow: 0 8px 24px rgba(15, 23, 42, 0.08); border-radius: 10px;',
      formatter: (params) => {
        const d = params?.data || {}
        const tip = d?.raw?.tip ? `<br/>${d.raw.tip}` : ''
        return `${d.name}<br/>投入：${d.effort}/10<br/>收益：${d.impact}/10<br/>ROI：${d.roi}/10${tip}`
      },
    },
    xAxis: {
      name: '投入（难度/时间）',
      nameTextStyle: { color: COLORS.textSub, fontSize: 11 },
      min: 0,
      max: 10,
      splitNumber: 5,
      axisLine: { lineStyle: { color: COLORS.gridLine } },
      axisLabel: { color: COLORS.textSub },
      splitLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.25)', type: 'dashed' } },
    },
    yAxis: {
      name: '收益（提升强度）',
      nameTextStyle: { color: COLORS.textSub, fontSize: 11 },
      min: 0,
      max: 10,
      splitNumber: 5,
      axisLine: { lineStyle: { color: COLORS.gridLine } },
      axisLabel: { color: COLORS.textSub },
      splitLine: { lineStyle: { color: 'rgba(148, 163, 184, 0.25)', type: 'dashed' } },
    },
    series: [
      {
        type: 'scatter',
        data,
        symbolSize: (val) => {
          const roi = Number(val?.[2] || 0)
          return 10 + roi * 4
        },
        itemStyle: {
          color: (params) => {
            const roi = params?.data?.roi ?? 0
            if (roi >= 8) return 'rgba(20, 184, 166, 0.85)'
            if (roi >= 6) return 'rgba(45, 212, 191, 0.75)'
            return 'rgba(153, 246, 228, 0.7)'
          },
          borderColor: 'rgba(255, 255, 255, 0.95)',
          borderWidth: 2,
          shadowBlur: 14,
          shadowColor: 'rgba(20, 184, 166, 0.12)',
        },
        emphasis: {
          scale: true,
        },
        label: {
          show: true,
          formatter: (p) => p.data?.raw?.short || p.data?.name,
          color: COLORS.textMain,
          fontWeight: 700,
          fontSize: 11,
          position: 'top',
        },
      },
    ],
  }
}

export const getMiniSankeyOption = ({ nodes = [], links = [], title = '' } = {}) => {
  return {
    ...getSankeyOption({ nodes, links, title }),
    series: [
      {
        type: 'sankey',
        left: '2%',
        right: '2%',
        top: title ? 34 : 6,
        bottom: 6,
        nodeGap: 10,
        nodeWidth: 10,
        emphasis: { focus: 'adjacency' },
        data: (Array.isArray(nodes) ? nodes : []).map((n) => ({
          ...n,
          itemStyle: {
            color: n.color || '#99f6e4',
            borderColor: 'rgba(255, 255, 255, 0.9)',
            borderWidth: 1,
          },
          label: {
            color: COLORS.textMain,
            fontSize: 11,
            fontWeight: 600,
          },
        })),
        links: (Array.isArray(links) ? links : []).map((l) => ({
          ...l,
          lineStyle: {
            color: 'rgba(20, 184, 166, 0.32)',
            opacity: 0.9,
            curveness: 0.45,
          },
        })),
        lineStyle: { color: 'source', curveness: 0.45 },
      },
    ],
  }
}

export const getFlowGraphOption = ({ nodes, links, title }) => {
  return {
    title: title
      ? {
        text: title,
        left: 0,
        top: 0,
        textStyle: { color: COLORS.textMain, fontSize: 14, fontWeight: 'bold' },
      }
      : undefined,
    tooltip: {
      trigger: 'item',
      backgroundColor: 'rgba(255, 255, 255, 0.92)',
      borderColor: COLORS.gridLine,
      borderWidth: 1,
      textStyle: { color: COLORS.textMain },
      extraCssText: 'box-shadow: 0 8px 24px rgba(15, 23, 42, 0.08); border-radius: 10px;',
      formatter: (params) => {
        const d = params?.data || {}
        const detail = Array.isArray(d.detail) ? d.detail : []
        if (!detail.length) return d.name
        return `${d.name}<br/>${detail.map((t) => `• ${t}`).join('<br/>')}`
      },
    },
    series: [
      {
        type: 'graph',
        layout: 'none',
        roam: true,
        data: nodes.map((n) => ({
          ...n,
          symbol: 'roundRect',
          symbolSize: n.symbolSize || [220, 54],
          itemStyle: {
            color: n.color || '#f0fdfa',
            borderColor: 'rgba(20, 184, 166, 0.35)',
            borderWidth: 1,
          },
          label: {
            show: true,
            color: COLORS.textMain,
            fontWeight: 700,
            fontSize: 12,
            overflow: 'truncate',
            width: 200,
          },
        })),
        links: links.map((l) => ({
          ...l,
          lineStyle: {
            color: 'rgba(20, 184, 166, 0.55)',
            width: 2,
          },
        })),
        edgeSymbol: ['none', 'arrow'],
        edgeSymbolSize: 10,
      },
    ],
  }
}
/**
 * [新增] 获取能力成长曲线配置 (折线面积图)
 * 用于展示：当前能力 vs 目标要求
 */
export const getGrowthLineOption = (categories, currentData, targetData) => {
  return {
    tooltip: {
      trigger: 'axis',
      backgroundColor: 'rgba(255, 255, 255, 0.95)',
      borderColor: COLORS.gridLine,
      textStyle: { color: COLORS.textMain },
      extraCssText: 'box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1); border-radius: 8px;'
    },
    legend: {
      data: ['当前能力', '目标要求'],
      bottom: 0,
      textStyle: { color: COLORS.textSub, fontSize: 12 }
    },
    grid: { left: '3%', right: '4%', bottom: '10%', top: '10%', containLabel: true },
    xAxis: {
      type: 'category',
      boundaryGap: false,
      data: categories,
      axisLine: { lineStyle: { color: COLORS.gridLine } },
      axisLabel: { color: COLORS.textSub, fontSize: 11, interval: 0, rotate: 0 }
    },
    yAxis: {
      type: 'value',
      max: 100,
      splitLine: { lineStyle: { color: COLORS.gridLine, type: 'dashed' } },
      axisLabel: { color: COLORS.textSub, fontSize: 11 }
    },
    series: [
      {
        name: '当前能力',
        type: 'line',
        smooth: true,
        data: currentData,
        itemStyle: { color: COLORS.primary },
        lineStyle: { width: 3 },
        areaStyle: {
          color: new echarts.graphic.LinearGradient(0, 0, 0, 1, [
            { offset: 0, color: 'rgba(20, 184, 166, 0.4)' },
            { offset: 1, color: 'rgba(20, 184, 166, 0.05)' }
          ])
        }
      },
      {
        name: '目标要求',
        type: 'line',
        smooth: true,
        data: targetData,
        itemStyle: { color: COLORS.warning },
        lineStyle: { type: 'dashed', width: 2, color: COLORS.warning },
        symbol: 'circle',
        symbolSize: 6
      }
    ]
  };
};
/**
 * [新增] 获取多维竞争力象限配置 (散点图)
 * 用于展示：用户在行业中的定位 (X:技术深度, Y:业务影响力)
 */
export const getCompetenceScatterOption = (dataPoints) => {
  return {
    tooltip: {
      formatter: (param) => {
        const d = param.data;
        return `<div style="padding:4px">
          <div style="font-weight:bold">${d.name}</div>
          <div>技术深度: ${d.value[0]}</div>
          <div>业务影响: ${d.value[1]}</div>
        </div>`;
      },
      backgroundColor: 'rgba(255, 255, 255, 0.95)',
      borderColor: COLORS.gridLine,
      textStyle: { color: COLORS.textMain }
    },
    xAxis: {
      name: '技术深度',
      nameLocation: 'middle',
      nameGap: 25,
      min: 0,
      max: 100,
      splitLine: { show: false },
      axisLine: { lineStyle: { color: COLORS.gridLine } },
      axisLabel: { color: COLORS.textSub }
    },
    yAxis: {
      name: '业务影响力',
      nameLocation: 'middle',
      nameGap: 30,
      min: 0,
      max: 100,
      splitLine: { show: false },
      axisLine: { lineStyle: { color: COLORS.gridLine } },
      axisLabel: { color: COLORS.textSub }
    },
    // 高亮“明星区” (右上角)
    markArea: {
      silent: true,
      itemStyle: { color: 'rgba(20, 184, 166, 0.08)' },
      label: { show: true, position: 'insideTopRight', formatter: '明星区', color: COLORS.primary, fontSize: 12 },
      data: [[{ xAxis: 75, yAxis: 75 }, { xAxis: 100, yAxis: 100 }]]
    },
    series: [{
      type: 'scatter',
      data: dataPoints,
      symbol: 'circle',
      symbolSize: (params) => params.name === '您' ? 16 : 10,
      itemStyle: {
        color: (params) => params.name === '您' ? COLORS.primary : '#cbd5e1',
        shadowBlur: 10,
        shadowColor: 'rgba(0,0,0,0.1)'
      },
      label: {
        show: true,
        formatter: '{b}',
        position: 'top',
        color: COLORS.textMain,
        fontWeight: 'bold',
        fontSize: 12
      }
    }]
  };
};
/**
 * [优化] 获取桑基图配置
 * @param {Object} params - 配置对象 { nodes, links, title }
 * 如果 nodes 未提供或无效，将自动从 links 中提取
 */
export const getSankeyOption = (params) => {
  // 解构参数，并设置默认值，防止 undefined
  const { nodes: inputNodes, links: inputLinks, title = '' } = params || {};

  let processedNodes = inputNodes;
  const links = Array.isArray(inputLinks) ? inputLinks : [];
  // 如果没有提供 nodes 数组，根据 links 自动生成
  if (!Array.isArray(processedNodes) || processedNodes.length === 0) {
    const nodeSet = new Set();
    // 安全遍历 links
    links.forEach(l => {
      if (l && l.source) nodeSet.add(l.source);
      if (l && l.target) nodeSet.add(l.target);
    });
    processedNodes = Array.from(nodeSet).map(name => ({ name }));
  }
  return {
    tooltip: {
      trigger: 'item',
      triggerOn: 'mousemove',
      backgroundColor: 'rgba(255, 255, 255, 0.95)',
      borderColor: COLORS.gridLine,
      textStyle: { color: COLORS.textMain }
    },
    series: [{
      type: 'sankey',
      layout: 'none',
      focusNodeAdjacency: 'allEdges',
      data: processedNodes.map(n => ({
        ...n,
        itemStyle: {
          color: n.color || (n.name.includes('终局') ? COLORS.warning :
            n.name.includes('当前') ? COLORS.secondary : COLORS.primary),
          opacity: 0.8
        },
        label: {
          color: COLORS.textMain,
          fontWeight: 600,
          fontSize: 12
        }
      })),
      links: links.map(l => ({
        ...l,
        lineStyle: {
          color: 'source',
          curveness: 0.5,
          opacity: 0.6
        }
      })),
      emphasis: { focus: 'adjacency' }
    }]
  };
};

/**
 * [优化] 获取风险预警仪表盘配置 (半圆，带颜色区段)
 * 0-40: 绿色 (低风险), 40-70: 黄色 (中风险), 70-100: 红色 (高风险)
 */
export const getRiskGaugeOption = (value, comment = '') => {
  const safeValue = Math.max(0, Math.min(100, Number(value) || 0));

  // 根据分数决定指针颜色和标题
  let statusColor = COLORS.success;
  let statusText = '低风险';
  if (safeValue > 70) { statusColor = COLORS.danger; statusText = '高风险'; }
  else if (safeValue > 40) { statusColor = COLORS.warning; statusText = '中风险'; }

  return {
    series: [{
      type: 'gauge',
      startAngle: 180,
      endAngle: 0,
      min: 0,
      max: 100,
      splitNumber: 5,
      radius: '90%',
      center: ['50%', '60%'],
      itemStyle: { color: statusColor },
      progress: {
        show: true,
        width: 18,
        roundCap: true
      },
      pointer: { show: false },
      axisLine: {
        lineStyle: {
          width: 18,
          color: [
            [0.4, COLORS.success],
            [0.7, COLORS.warning],
            [1, COLORS.danger]
          ]
        }
      },
      axisTick: { show: false },
      splitLine: { show: false },
      axisLabel: { show: false },
      detail: {
        valueAnimation: true,
        offsetCenter: [0, '10%'],
        fontSize: 24,
        fontWeight: 'bold',
        color: COLORS.textMain,
        formatter: '{value}',
        backgroundColor: 'transparent'
      },
      title: {
        show: true,
        offsetCenter: [0, '40%'],
        fontSize: 14,
        color: COLORS.textSub,
        formatter: statusText
      },
      data: [{ value: safeValue, name: statusText }]
    }],
    // 底部文字说明
    graphic: comment ? [{
      type: 'text',
      left: 'center',
      bottom: '5%',
      style: {
        text: comment,
        fill: COLORS.textSub,
        font: '12px sans-serif',
        textAlign: 'center'
      }
    }] : []
  };
};
