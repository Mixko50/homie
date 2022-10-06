//
//  States.swift
//  Homie
//
//  Created by Mixko on 10/5/22.
//

import SwiftUI

struct States: View {
    var body: some View {
        NavigationView {
            List {
                StateBox()
                StateBox()
                StateBox()
                StateBox()
                StateBox()
                StateBox()
            }.navigationTitle("All States")
        }
    }
}

struct States_Previews: PreviewProvider {
    static var previews: some View {
        States()
    }
}
