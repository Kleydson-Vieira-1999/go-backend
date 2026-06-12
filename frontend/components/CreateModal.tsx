'use client'

import { useRouter } from 'next/navigation'

interface CreateModalProps {
  isOpen: boolean
  onClose: () => void
  title: string
  description: string
  confirmUrl: string
}

export function CreateModal({ isOpen, onClose, title, description, confirmUrl }: CreateModalProps) {
  const router = useRouter()

  if (!isOpen) return null

  return (
    <div className="fixed inset-0 bg-black/60 z-50 flex items-center justify-center p-4 backdrop-blur-sm">
      <div className="bg-white dark:bg-zinc-900 p-6 rounded-xl border border-zinc-200 dark:border-zinc-800 max-w-md w-full shadow-xl">
        <h3 className="font-bold text-lg mb-2">{title}</h3>
        <p className="text-sm text-zinc-500 mb-6">{description}</p>
        <div className="flex justify-end gap-3 border-t border-zinc-200 dark:border-zinc-800 pt-4">
          <button 
            onClick={onClose} 
            className="text-sm font-medium text-zinc-500 hover:text-zinc-700"
          >
            Cancelar
          </button>
          <button
            onClick={() => router.push(confirmUrl)}
            className="bg-emerald-600 hover:bg-emerald-700 text-white px-4 py-2 rounded-lg text-sm font-medium transition-colors shadow-sm"
          >
            Ir para Formulário
          </button>
        </div>
      </div>
    </div>
  )
}