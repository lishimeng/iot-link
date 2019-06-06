<template>
    <div class="body">
      <div class="search">
        <el-row :gutter="15">
          <search-option />
        </el-row>
      </div>
      <div>
        <el-row>
          <el-tabs v-model="activeName">
            <el-tab-pane label="精神药品" name="first"><items :list="sample" /></el-tab-pane>
            <el-tab-pane label="麻醉药品" name="second"><items :list="sample" /></el-tab-pane>
            <el-tab-pane label="非药用类麻醉药品和精神药品" name="third"><items :list="sample" /></el-tab-pane>
            <el-tab-pane label="易制毒化学品" name="fourth"><items :list="sample" /></el-tab-pane>
          </el-tabs>
        </el-row>
      </div>
    </div>
</template>

<script>
import SearchOption from './components/SearchOption'
import Items from './components/Items'
export default {
  name: "CasList",
  components: { SearchOption, Items },
  data: () => {
    return{
      sample: [
        { id: 1, cn: '醋托啡', en: 'Acetorphine', cas: '25333-77-1', formula: 'C19H21NO3' },
        { id: 2, cn: '乙酰阿法甲基芬太尼', en: 'Acetyl-alpha-methylfentanyl ', cas: '101860-00-8', formula: 'C22H28N2O' },
        { id: 3, cn: '醋美沙多', en: 'Acetylmethadol ', cas: '509-74-0', formula: 'C23H31NO2' },
        { id: 4, cn: '阿芬太尼', en: 'Alfentanil', cas: '71195-58-9', formula: 'C21H32N6O3' },
        { id: 5, cn: '烯丙罗定', en: 'Allylprodine ', cas: '25384-17-2', formula: 'C18H25NO2' },
        { id: 6, cn: '阿醋美沙多', en: 'Alphacetylmethadol ', cas: '17199-58-5', formula: 'C17H25NO2' },
      ]
    }
}
}
</script>

<style scoped>
  .body {
    margin: 15px;
  }
  .search {
    margin-top: 15px;
    margin-bottom: 15px;
  }

</style>
