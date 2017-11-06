<template>
  <ve-line :data="chartData" :settings="chartSettings"></ve-line>
</template>

<script>
  export default {
    name: 'PlatFansChart',
    data () {
      return {
        loading: false,
        chartData: {
          columns: ['date', 'balance', 'age'],
          rows: [
            {'date': '1月1日', 'balance': 123, 'age': 3},
            {'date': '1月2日', 'balance': 1223, 'age': 6},
            {'date': '1月3日', 'balance': 2123, 'age': 9},
            {'date': '1月4日', 'balance': 4123, 'age': 12},
            {'date': '1月5日', 'balance': 3123, 'age': 15},
            {'date': '1月6日', 'balance': 7123, 'age': 20}
          ]
        },
        chartSettings: {
          labelMap: {
            balance: '余额',
            age: '年龄'
          }
        }
      }
    },
    created () {
      this.fetchList()
    },
    methods: {
      // net io
      fetchList () {
        this.$http.get(this.url, {
//          params: {
//            ID: 12345
//          }
        }).then((response) => {
          this.chartData.rows = response.data
        }).then()
      },
      reload () {
        this.fetchList()
      }
    }
  }
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }
</style>
