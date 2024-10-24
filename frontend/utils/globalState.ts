import ItemType from '@/types/item';
import { atom } from 'jotai';

const sampleItem:ItemType = {
    itemName:"メニュー",
    itemPrice:200
}

const itemAtom = atom<ItemType[]>([sampleItem])

export default itemAtom