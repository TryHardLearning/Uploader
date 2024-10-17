package com.paysystem.service.CRUD;

import com.paysystem.model.DTO.PaymentsDTO;
import com.paysystem.model.Payments;
import com.paysystem.repository.PaymentsRepository;
import org.modelmapper.ModelMapper;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class PaymentsCRUD {
    private PaymentsRepository repository;

    @Autowired
    public PaymentsCRUD(PaymentsRepository rep){
        this.repository = rep;
    }
    private ModelMapper model;

    public List<PaymentsDTO> listAll(){
        return (List<PaymentsDTO>) repository.findAll().stream().map(pay -> PaymentsDTO.builder().price(pay.getPrice()).client(pay.getClient()).date(pay.getDate()).status(pay.getStatus()).build());

    }
    public PaymentsDTO newPayment(Payments pay){
        return PaymentsDTO.builder().price(pay.getPrice()).client(pay.getClient()).date(pay.getDate()).status(pay.getStatus()).build();
    }
}
