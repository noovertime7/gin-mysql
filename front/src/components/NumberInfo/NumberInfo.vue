<template>
  <div :class="[prefixCls]">
    <slot name="subtitle">
      <div :class="[`${prefixCls}-subtitle`]">{{ typeof subTitle === 'string' ? subTitle : subTitle() }}</div>
    </slot>
    <div class="number-info-value">
      <span>{{ total }}</span>
      <span class="sub-total">
        {{ message }}
      </span>
      <span class="sub-total">&ensp;</span>
      <span class="sub-total">
        {{ subTotal }}
        <icon :type="`caret-${status}`" />
      </span>
      <span class="sub-total">&ensp;</span>
      <span class="sub-total">
        {{ decreaseTotal }}
        <icon :type="`caret-up`" />
      </span>
    </div>
  </div>
</template>

<script>
import Icon from 'ant-design-vue/es/icon'

export default {
  name: 'NumberInfo',
  props: {
    prefixCls: {
      type: String,
      default: 'ant-pro-number-info'
    },
    total: {
      type: Number,
      default: 0
    },
    subTotal: {
      type: Number,
      default: 0
    },
    decreaseTotal: {
      type: Number,
      default: 100
    },
    subTitle: {
      type: [String, Function],
      default: ''
    },
    status: {
      type: String,
      default: 'up'
    },
    message: {
      type: String,
      default: 'unknown'
    }
  },
  components: {
    Icon
  },
  data () {
    return {}
  }
}
</script>

<style lang="less" scoped>
  @import "index";
</style>
