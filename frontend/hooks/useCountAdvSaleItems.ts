import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'

const useCountAdvSaleItems = () => {
  const [cartItems] = useAtom(cartItemsAtom)
  //TODO: 現状(前)が文字として入っているものを前売り券としてカウント。後に修正する。
  const advSaleItems = cartItems.filter(cartItems =>
    cartItems.name.includes('前'),
  )

  return advSaleItems
}

export default useCountAdvSaleItems
