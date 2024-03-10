export function isApp() {
  const isMobile = /Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(
    navigator.userAgent
  )
  const isSmallScreen = window.innerWidth < 768
  console.log('isMobile', isMobile)
  console.log('isSmallScreen', isSmallScreen)
  return isMobile || isSmallScreen
}
