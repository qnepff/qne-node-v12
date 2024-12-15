<template>
  <div class="toggle-switch">
    <input
      type="checkbox"
      :id="id"
      :checked="modelValue"
      @change="$emit('update:modelValue', $event.target.checked)"
      class="toggle-switch-checkbox"
    >
    <label class="toggle-switch-label" :for="id">
      <span class="toggle-switch-inner" :data-yes="labelOn" :data-no="labelOff"></span>
      <span class="toggle-switch-switch"></span>
    </label>
  </div>
</template>

<script setup>
defineProps({
  modelValue: Boolean,
  labelOn: {
    type: String,
    default: 'On'
  },
  labelOff: {
    type: String,
    default: 'Off'
  },
  id: {
    type: String,
    required: true
  }
});

defineEmits(['update:modelValue']);
</script>

<style scoped>
.toggle-switch {
  position: relative;
  width: 75px;
  display: inline-block;
  vertical-align: middle;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  text-align: left;
}
.toggle-switch-checkbox {
  display: none;
}
.toggle-switch-label {
  display: block;
  overflow: hidden;
  cursor: pointer;
  border: 0 solid #bbb;
  border-radius: 20px;
  margin: 0;
}
.toggle-switch-inner {
  display: block;
  width: 200%;
  margin-left: -100%;
  transition: margin 0.3s ease-in 0s;
}
.toggle-switch-inner:before,
.toggle-switch-inner:after {
  display: block;
  float: left;
  width: 50%;
  height: 34px;
  padding: 0;
  line-height: 34px;
  font-size: 14px;
  color: white;
  font-weight: bold;
  box-sizing: border-box;
}
.toggle-switch-inner:before {
  content: attr(data-yes);
  text-transform: uppercase;
  padding-left: 10px;
  background-color: #2E5E20;
  color: #fff;
}
.toggle-switch-inner:after {
  content: attr(data-no);
  text-transform: uppercase;
  padding-right: 10px;
  background-color: #2E5E20;
  color: #fff;
  text-align: right;
}
.toggle-switch-switch {
  display: block;
  width: 24px;
  margin: 5px;
  background: #fff;
  position: absolute;
  top: 0;
  bottom: 0;
  right: 40px;
  border: 0 solid #bbb;
  border-radius: 20px;
  transition: all 0.3s ease-in 0s;
}
.toggle-switch-checkbox:checked + .toggle-switch-label .toggle-switch-inner {
  margin-left: 0;
}
.toggle-switch-checkbox:checked + .toggle-switch-label .toggle-switch-switch {
  right: 0px;
}
</style>
