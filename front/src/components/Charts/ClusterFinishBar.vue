<template>
  <div :style="{ padding: '0 0 32px 32px' }">
    <h4 :style="{ marginBottom: '20px' }">{{ title }}</h4>
    <v-chart
      :scale="clusterTaskDataByDateScale"
      height="254"
      :data="data"
      :forceFit="true"
      :padding="['auto', 'auto', '20', '50']">
      <v-tooltip />
      <v-axis />
      <v-bar position="date*finish_num"/>
    </v-chart>
  </div>
</template>

<script>
export default {
  name: 'ClusterFinishBar',
  props: {
    title: {
      type: String,
      default: ''
    },
    data: {
      type: Array,
      default: () => {
        return []
      }
    },
    tooltip: {
      type: Array,
      default: () => {
        return [
          'x*y',
          (x, y) => ({
            name: x,
            value: y
          })
        ]
      }
    }
  },
  data () {
    return {
      clusterTaskDataByDateScale: [
        {
          dataKey: 'date',
          alias: '日期'
        },
        {
          dataKey: 'finish_num',
          alias: '完成数'
        }]
    }
  }
}
</script>
