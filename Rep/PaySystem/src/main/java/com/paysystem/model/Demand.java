package com.paysystem.model;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.NoArgsConstructor;
import org.hibernate.annotations.Type;

import javax.persistence.*;
import java.util.List;

@Entity @NoArgsConstructor
@AllArgsConstructor
@Builder
public class Demand {
    @Id @GeneratedValue(strategy = GenerationType.AUTO) @Type(type = "long")
    private Long id;
    @OneToMany
    private List<Product> productList;
}
