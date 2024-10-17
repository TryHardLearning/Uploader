package com.paysystem.model.DTO;

import com.paysystem.model.Client;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;

import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.OneToOne;
import java.math.BigDecimal;
import java.time.LocalDateTime;

@Entity @NoArgsConstructor @AllArgsConstructor @Builder
public class PaymentsDTO {
    @Id
    private Long id ;

    private BigDecimal price;
    private LocalDateTime date;
    @OneToOne
    private Client client;
    private String status;
}
