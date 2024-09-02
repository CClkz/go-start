
/**
 * 求平方根
 * @param {*} x 
 */
function sqrt(x){
  if(x===0||x===1){
    return x
  }

  let res = x
  const epsilon = 0.000001; // 精度
  while(Math.abs(res*res - x)>epsilon){
    res = (res + x/res)/2
  }
  return res
}