package com.paysystem.controller;

import com.paysystem.model.DTO.PaymentsDTO;
import com.paysystem.service.CRUD.PaymentsCRUD;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.util.UriBuilder;

import javax.validation.constraints.NotNull;
import java.net.URI;
import java.net.URL;
import java.util.List;

@RestController
@RequestMapping("/payments")
public class PaymentsController {

    private PaymentsCRUD PayCRUD;

    @Autowired
    public PaymentsController(PaymentsCRUD pay){
        PayCRUD = pay;
    }

    @GetMapping
    public List<PaymentsDTO> list (){
        return PayCRUD.listAll();
    }
    @DeleteMapping("/{id}")
    public ResponseEntity<PaymentsDTO> details(@PathVariable @NotNull Long id, UriBuilder uriBuilder){
            PaymentsDTO pay = PayCRUD.newPayment();
                    URI address = uriBuilder.path("payments/{id}");
    }
}
