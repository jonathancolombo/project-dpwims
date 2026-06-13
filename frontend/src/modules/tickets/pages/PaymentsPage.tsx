import { useEffect, useState } from "react";
import MainLayout from "../../../core/layout/MainLayout";
import { getPayments, deletePayment } from "../api/payments_api";
import { useNavigate } from "react-router-dom";
import type { Payment } from "../types/payment";

export default function PaymentsPage() {
    const [payments, setPayments] = useState<Payment[]>([]);
    const [loading, setLoading] = useState(true);
    const [message, setMessage] = useState("");
    const navigate = useNavigate();

    useEffect(() => {
        getPayments()
            .then(response => {
                setPayments(response.data);
                setLoading(false);
            })
            .catch(() => {
                setMessage("Errore nel caricamento dei pagamenti.");
                setLoading(false);
            });
    }, []);

    const handleDelete = async (uuid: string) => {
        if (!confirm("Sei sicuro di voler eliminare questo pagamento?")) return;

        try {
            await deletePayment(uuid);
            setPayments(payments => payments.filter(payment => payment.uuid !== uuid));
        } catch {
            setMessage("Errore durante l'eliminazione del pagamento.");
        }
    };

    const paymentMethodLabels: Record<Payment["payment_method"], string> = {
        credit_card: "Carta di credito",
        banknotes: "Contanti",
        bank_transfer: "Bonifico bancario"
    };

    if (loading) {
        return (
            <MainLayout>
                <div className="p-6 text-center text-gray-600">Caricamento...</div>
            </MainLayout>
        );
    }

    return (
        <MainLayout>
            <div className="p-6 space-y-6">
                <div className="flex justify-between items-center">
                    <h1 className="text-3xl font-bold">Pagamenti</h1>

                    <button
                        onClick={() => navigate("/payments/create")}
                        className="bg-blue-600 text-white px-4 py-2 rounded-lg hover:bg-blue-700"
                    >
                        + Nuovo pagamento
                    </button>
                </div>

                {message && (
                    <div className="p-3 bg-red-100 text-red-700 rounded-lg">
                        {message}
                    </div>
                )}

                <div className="space-y-4">
                    {payments.map(payment => (
                        <div
                            key={payment.uuid}
                            className="p-4 bg-white shadow rounded-lg border space-y-4"
                        >
                            {/* Header */}
                            <div className="flex justify-between items-center">
                                <h2 className="text-lg font-semibold">
                                    Pagamento #{payment.uuid}
                                </h2>

                                <span className="px-2 py-1 rounded bg-gray-200 text-gray-700">
                                    {paymentMethodLabels[payment.payment_method]}
                                </span>
                            </div>

                            {/* Info */}
                            <div className="text-gray-700 space-y-1">
                                <div><strong>Ticket ID:</strong> {payment.ticket_id}</div>
                                <div><strong>Importo:</strong> € {payment.amount}</div>
                                <div><strong>Riferimento provider:</strong> {payment.provider_reference}</div>
                            </div>

                            {/* Actions */}
                            <div className="flex gap-3 pt-2">
                                <button
                                    onClick={() => navigate(`/payments/${payment.uuid}`)}
                                    className="flex-1 bg-gray-100 hover:bg-gray-200 text-gray-700 py-2 rounded-lg"
                                >
                                    Modifica
                                </button>

                                <button
                                    onClick={() => handleDelete(payment.uuid)}
                                    className="flex-1 bg-red-600 hover:bg-red-700 text-white py-2 rounded-lg"
                                >
                                    Elimina
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            </div>
        </MainLayout>
    );
}
