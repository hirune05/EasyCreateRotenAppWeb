import { cartItemsAtom } from '@/utils/globalState'
import { useAtom } from 'jotai'

const useCountAdvSaleItems = () => {
  const [cartItems] = useAtom(cartItemsAtom)
  const advSaleItems = cartItems.filter(cartItems =>
    cartItems.name.includes('前売り券'),
  )

  return advSaleItems
}

export default useCountAdvSaleItems
