import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'

const useCountTotalPrice = () => {
  const [cartItems] = useAtom(cartItemsAtom)

  // '前売り券' を含まないアイテムをフィルタリング
  const usuSaleItems = cartItems.filter(item => !item.name.includes('前売り券'))

  // 合計金額を計算
  const totalPrice = usuSaleItems.reduce(
    (sum, item) => sum + (item.price ?? 0),
    0,
  )

  return totalPrice
}

export default useCountTotalPrice
