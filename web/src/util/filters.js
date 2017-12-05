export function host (url) {
  const host = url.replace(/^https?:\/\//, '').replace(/\/.*$/, '')
  const parts = host.split('.').slice(-3)
  if (parts[0] === 'www') parts.shift()
  return parts.join('.')
}

export function timeAgo (time) {
  const between = Date.now() / 1000 - Number(time)
  if (between < 3600) {
    return pluralize(~~(between / 60), ' minute')
  } else if (between < 86400) {
    return pluralize(~~(between / 3600), ' hour')
  } else {
    return pluralize(~~(between / 86400), ' day')
  }
}

export function timeToDay (timestamp) {
  let a = new Date(timestamp * 1000)
  let year = a.getFullYear()
  let month = a.getMonth() + 1
  let date = a.getDate()
  // let hour = a.getHours()
  // let min = a.getMinutes()
  // let sec = a.getSeconds()
  return year + '年 ' + month + '月 ' + date + '日 '
}

function pluralize (time, label) {
  if (time === 1) {
    return time + label
  }
  return time + label + 's'
}

export function shortTitle (val) {
  if (val.length <= 10) {
    return val
  } else {
    return val.slice(0, 10) + '...'
  }
}

export function shortText (val) {
  if (val.length <= 18) {
    return val
  } else {
    return val.slice(0, 17) + '...'
  }
}

export function talentStatus (val) {
  if (val) {
    return '已推'
  } else {
    return '未推'
  }
}
